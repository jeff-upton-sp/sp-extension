package main

import (
	"context"

	"github.com/jeff-upton-sp/sp-extension/internal/infra"
	"github.com/sailpoint/atlas-go/atlas/log"
)

func main() {
	ctx := context.Background()

	service, err := infra.NewExtensionService(ctx)
	if err != nil {
		log.Fatalf(ctx, "error: %v", err)
	}
	defer service.Close(ctx)

	if err := service.Run(ctx); err != nil {
		log.Errorf(ctx, "error: %v", err)
	}
}
