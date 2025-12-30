package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type InvokeInput struct {
}

func (input InvokeInput) Validate() error {
	return nil
}

type InvokeOutput struct {
	Result json.RawMessage `json:"result"`
}

func invoke(ctx context.Context, input InvokeInput, repo model.FunctionRepo) (InvokeOutput, error) {
	if err := input.Validate(); err != nil {
		return InvokeOutput{}, err
	}

	return InvokeOutput{}, fmt.Errorf("not implemented")
}
