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

func invoke(ctx context.Context, input InvokeInput, repo model.FunctionRepo, evaluator model.FunctionEvaluator) (InvokeOutput, error) {
	if err := input.Validate(); err != nil {
		return InvokeOutput{}, err
	}

	f, err := repo.FindByID(ctx, input.FunctionID)
	if err != nil {
		return InvokeOutput{}, err
	}

	result, err := evaluator.EvaluateFunction(ctx, f.SourceCode, input.Input)
	if err != nil {
		return InvokeOutput{}, fmt.Errorf("invoke '%s': %w", f.ID, err)
	}

	return InvokeOutput{
		Result: result,
	}, nil
}
