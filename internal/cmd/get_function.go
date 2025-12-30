package cmd

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type GetFunctionInput struct {
	FunctionID model.FunctionID `json:"functionId"`
}

func (input GetFunctionInput) Validate() error {
	if input.FunctionID == "" {
		return fmt.Errorf("FunctionID is required")
	}

	return nil
}

type GetFunctionOutput struct {
	Function model.Function `json:"function"`
}

func getFunction(ctx context.Context, input GetFunctionInput, repo model.FunctionRepo) (GetFunctionOutput, error) {
	if err := input.Validate(); err != nil {
		return GetFunctionOutput{}, err
	}

	f, err := repo.FindByID(ctx, input.FunctionID)
	if err != nil {
		return GetFunctionOutput{}, fmt.Errorf("get function '%s': %w", input.FunctionID, err)
	}

	return GetFunctionOutput{
		Function: f,
	}, nil
}
