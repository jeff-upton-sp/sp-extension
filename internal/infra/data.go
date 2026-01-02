package infra

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/jeff-upton-sp/sp-extension/internal/cmd"
)

//go:embed data/functions
var functionFS embed.FS

func (s *ExtensionService) loadFunctions(ctx context.Context) error {
	if err := fs.WalkDir(functionFS, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || !strings.HasSuffix(path, ".js") {
			return nil
		}

		sourceCode, err := functionFS.ReadFile(path)
		if err != nil {
			return err
		}

		functionName := filepath.Base(path)
		functionName = strings.TrimRight(functionName, "."+filepath.Ext(functionName))

		createFunctionInput := cmd.CreateFunctionInput{
			Name:       functionName,
			SourceCode: string(sourceCode),
		}

		if _, err := s.app.CreateFunction(ctx, createFunctionInput); err != nil {
			return fmt.Errorf("create function '%s': %w", path, err)
		}

		return nil
	}); err != nil {
		return nil
	}

	return nil
}

func (s *ExtensionService) loadData(ctx context.Context) error {
	if err := s.loadFunctions(ctx); err != nil {
		return err
	}

	return nil
}
