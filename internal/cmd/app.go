package cmd

import (
	"context"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type App struct {
	FunctionRepo      model.FunctionRepo
	FunctionEvaluator model.FunctionEvaluator
}

func (app *App) GetFunction(ctx context.Context, input GetFunctionInput) (GetFunctionOutput, error) {
	return getFunction(ctx, input, app.FunctionRepo)
}

func (app *App) Invoke(ctx context.Context, input InvokeInput) (InvokeOutput, error) {
	return invoke(ctx, input, app.FunctionRepo, app.FunctionEvaluator)
}
