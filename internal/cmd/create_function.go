package cmd

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
	"github.com/sailpoint/atlas-go/atlas/log"
)

type CreateFunctionInput struct {
	FunctionTypeID model.FunctionTypeID `json:"functionTypeId"`
	Name           string               `json:"name"`
	SourceCode     string               `json:"sourceCode"`
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

func createFunction(ctx context.Context, input CreateFunctionInput, repo model.FunctionRepo, typeRepo model.FunctionTypeRepo) (CreateFunctionOutput, error) {
	if err := input.Validate(); err != nil {
		return CreateFunctionOutput{}, err
	}

	if _, err := typeRepo.FindByID(ctx, input.FunctionTypeID); err != nil {
		return CreateFunctionOutput{}, fmt.Errorf("validate function type: %w", err)
	}

	f := model.Function{
		ID:             model.FunctionID(input.Name),
		FunctionTypeID: model.FunctionTypeID(input.FunctionTypeID),
		Name:           input.Name,
		SourceCode:     input.SourceCode,
	}

	if err := repo.Save(ctx, &f); err != nil {
		return CreateFunctionOutput{}, err
	}

	log.Infof(ctx, "created function %s - %s", f.ID, f.Name)

	return CreateFunctionOutput{
		Function: f,
	}, nil
}
