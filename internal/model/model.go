package model

import (
	"context"
	"encoding/json"
)

type FunctionID string

type Function struct {
	ID         FunctionID `json:"id"`
	Name       string     `json:"name"`
	SourceCode string     `json:"sourceCode"`
}

type FunctionRepo interface {
	FindAll(ctx context.Context) ([]Function, error)
	FindByID(ctx context.Context, id FunctionID) (Function, error)
}

type FunctionEvaluator interface {
	EvaluateFunction(ctx context.Context, sourceCode string, input json.RawMessage) (json.RawMessage, error)
}
