package model

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type FunctionID string

type Function struct {
	ID         FunctionID `json:"id"`
	Name       string     `json:"name"`
	SourceCode string     `json:"sourceCode"`
}

type FunctionRepo interface {
	FindByID(ctx context.Context, id FunctionID) (Function, error)
	Save(ctx context.Context, f *Function) error
}

type FunctionEvaluator interface {
	EvaluateFunction(ctx context.Context, sourceCode string, input json.RawMessage) (json.RawMessage, error)
}

func NewFunctionID() FunctionID {
	return FunctionID(uuid.NewString())
}
