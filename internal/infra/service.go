package infra

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/cmd"
	"github.com/jeff-upton-sp/sp-extension/internal/infra/memory"
	"github.com/sailpoint/atlas-go/atlas/application"
	"golang.org/x/sync/errgroup"
)

type ExtensionService struct {
	*application.Application
	app *cmd.App
}

func NewExtensionService(ctx context.Context) (*ExtensionService, error) {
	application, err := application.New("sp-policy", WithNullEventPublisher(), application.WithDefaultCodecRegistry())
	if err != nil {
		return nil, fmt.Errorf("create application: %w", err)
	}

	_, err = loadConfig(application.Config)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	functionRepo, err := memory.NewFunctionRepo()
	if err != nil {
		return nil, err
	}

	functionEvaluator := newGojaFunctionEvaluator()

	app := &cmd.App{
		FunctionRepo:      functionRepo,
		FunctionEvaluator: functionEvaluator,
	}

	s := &ExtensionService{}
	s.Application = application
	s.app = app

	if err := s.loadData(ctx); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *ExtensionService) Run(ctx context.Context) error {
	ctx, done := context.WithCancel(ctx)
	defer done()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return s.StartMetricsServer(ctx) })
	g.Go(func() error { return s.StartWebServer(ctx, s.buildRoutes()) })
	g.Go(func() error { return s.WaitForInterrupt(ctx, done) })

	if err := g.Wait(); err != nil && err != context.Canceled {
		return err
	}

	return nil
}

func (s *ExtensionService) Close(ctx context.Context) {
	if s.Application != nil {
		s.Application.Close()
	}
}
