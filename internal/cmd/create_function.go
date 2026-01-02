package cmd

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
	"github.com/sailpoint/atlas-go/atlas/log"
)

type CreateFunctionInput struct {
	Name       string `json:"name"`
	SourceCode string `json:"sourceCode"`
}

func (input CreateFunctionInput) Validate() error {
	if input.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if input.SourceCode == "" {
		return fmt.Errorf("SourceCode is required")
	}

	return nil
}

type CreateFunctionOutput struct {
	Function model.Function `json:"function"`
}

func createFunction(ctx context.Context, input CreateFunctionInput, repo model.FunctionRepo) (CreateFunctionOutput, error) {
	if err := input.Validate(); err != nil {
		return CreateFunctionOutput{}, err
	}

	f := model.Function{
		Name:       input.Name,
		SourceCode: input.SourceCode,
	}

	if err := repo.Save(ctx, &f); err != nil {
		return CreateFunctionOutput{}, err
	}

	log.Infof(ctx, "created function %s - %s", f.ID, f.Name)

	return CreateFunctionOutput{
		Function: f,
	}, nil
}
