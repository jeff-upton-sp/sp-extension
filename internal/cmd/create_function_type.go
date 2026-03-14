package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
	"github.com/sailpoint/atlas-go/atlas/log"
)

type CreateFunctionTypeInput struct {
	ID           model.FunctionTypeID `json:"id"`
	InputSchema  json.RawMessage      `json:"inputSchema"`
	OutputSchema json.RawMessage      `json:"outputSchema"`
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

func createFunctionType(ctx context.Context, input CreateFunctionTypeInput, repo model.FunctionTypeRepo) (CreateFunctionTypeOutput, error) {
	if err := input.Validate(); err != nil {
		return CreateFunctionTypeOutput{}, err
	}

	log.Infof(ctx, "created function type %s", input.ID)
	return CreateFunctionTypeOutput{}, fmt.Errorf("not implemented")
}
