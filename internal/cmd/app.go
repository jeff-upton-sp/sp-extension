package cmd

import (
	"context"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type App struct {
	FunctionRepo      model.FunctionRepo
	FunctionEvaluator model.FunctionEvaluator
	FunctionTypeRepo  model.FunctionTypeRepo
}

func (app *App) CreateFunction(ctx context.Context, input CreateFunctionInput) (CreateFunctionOutput, error) {
	return createFunction(ctx, input, app.FunctionRepo)
}

func (app *App) CreateFunctionType(ctx context.Context, input CreateFunctionTypeInput) (CreateFunctionTypeOutput, error) {
	return createFunctionType(ctx, input, app.FunctionTypeRepo)
}

func (app *App) GetFunction(ctx context.Context, input GetFunctionInput) (GetFunctionOutput, error) {
	return getFunction(ctx, input, app.FunctionRepo)
}

func (app *App) Invoke(ctx context.Context, input InvokeInput) (InvokeOutput, error) {
	return invoke(ctx, input, app.FunctionRepo, app.FunctionEvaluator)
}
