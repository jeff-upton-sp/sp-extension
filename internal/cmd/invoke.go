package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type InvokeInput struct {
	FunctionID model.FunctionID `json:"functionId"`
	Input      json.RawMessage  `json:"input"`
}

func (input InvokeInput) Validate() error {
	if input.FunctionID == "" {
		return fmt.Errorf("FunctionID is required")
	}

	return nil
}

type InvokeOutput struct {
	Result json.RawMessage `json:"result"`
}

type FunctionProvider interface {
	FindByID(ctx context.Context, id model.FunctionID) (model.Function, error)
}

func invoke(ctx context.Context, input InvokeInput, provider FunctionProvider, typeRepo model.FunctionTypeRepo, evaluator model.FunctionEvaluator, schemaCompiler model.SchemaCompiler) (InvokeOutput, error) {
	if err := input.Validate(); err != nil {
		return InvokeOutput{}, err
	}

	f, err := provider.FindByID(ctx, input.FunctionID)
	if err != nil {
		return InvokeOutput{}, err
	}

	ft, err := typeRepo.FindByID(ctx, f.FunctionTypeID)
	if err != nil {
		return InvokeOutput{}, err
	}

	// TODO: should not compile every time...
	inputSchema, err := schemaCompiler.CompileSchema(ctx, ft.InputSchema)
	if err != nil {
		return InvokeOutput{}, fmt.Errorf("compile input schema: %w", err)
	}

	outputSchema, err := schemaCompiler.CompileSchema(ctx, ft.OutputSchema)
	if err != nil {
		return InvokeOutput{}, fmt.Errorf("compile output schema: %w", err)
	}

	if err := inputSchema.ValidateJSON(ctx, []byte(input.Input)); err != nil {
		return InvokeOutput{}, fmt.Errorf("validate input: %w", err)
	}

	result, err := evaluator.EvaluateFunction(ctx, f.SourceCode, input.Input)
	if err != nil {
		return InvokeOutput{}, fmt.Errorf("invoke '%s': %w", f.ID, err)
	}

	if err := outputSchema.ValidateJSON(ctx, []byte(result)); err != nil {
		return InvokeOutput{}, fmt.Errorf("validate output: %w", err)
	}

	return InvokeOutput{
		Result: result,
	}, nil
}
