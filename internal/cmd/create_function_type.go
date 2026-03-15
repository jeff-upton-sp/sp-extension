package cmd

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
	"github.com/sailpoint/atlas-go/atlas/log"
)

type CreateFunctionTypeInput struct {
	ID           model.FunctionTypeID `json:"id"`
	InputSchema  model.RawSchema      `json:"inputSchema"`
	OutputSchema model.RawSchema      `json:"outputSchema"`
}

func (input CreateFunctionTypeInput) Validate() error {
	if input.ID == "" {
		return fmt.Errorf("ID is required")
	}

	if input.InputSchema == nil {
		return fmt.Errorf("InputSchema is required")
	}

	if len(input.OutputSchema) == 0 {
		return fmt.Errorf("OutputSchema is required")
	}

	return nil
}

type CreateFunctionTypeOutput struct {
	FunctionType model.FunctionType `json:"functionType"`
}

func createFunctionType(ctx context.Context, input CreateFunctionTypeInput, repo model.FunctionTypeRepo, schemaCompiler model.SchemaCompiler) (CreateFunctionTypeOutput, error) {
	if err := input.Validate(); err != nil {
		return CreateFunctionTypeOutput{}, err
	}

	if _, err := schemaCompiler.CompileSchema(ctx, input.InputSchema); err != nil {
		return CreateFunctionTypeOutput{}, fmt.Errorf("compile input schema: %w", err)
	}

	if _, err := schemaCompiler.CompileSchema(ctx, input.OutputSchema); err != nil {
		return CreateFunctionTypeOutput{}, fmt.Errorf("compile output schema: %w", err)
	}

	functionType := model.FunctionType{
		ID:           input.ID,
		InputSchema:  input.InputSchema,
		OutputSchema: input.OutputSchema,
	}

	if err := repo.Save(ctx, functionType); err != nil {
		return CreateFunctionTypeOutput{}, err
	}

	log.Infof(ctx, "created function type %s", input.ID)

	return CreateFunctionTypeOutput{
		FunctionType: functionType,
	}, nil
}
