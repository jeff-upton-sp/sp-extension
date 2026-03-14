package model

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type FunctionID string

type FunctionTypeID string

type Schema interface {
	ValidateJSON(ctx context.Context, data []byte) error
}

type FunctionType struct {
	ID           FunctionTypeID `json:"id"`
	InputSchema  Schema         `json:"inputSchema"`
	OutputSchema Schema         `json:"outputSchema"`
}

type Function struct {
	ID         FunctionID `json:"id"`
	Name       string     `json:"name"`
	SourceCode string     `json:"sourceCode"`
}

type FunctionProvider interface {
	FindByID(ctx context.Context, id FunctionID) (Function, error)
}

type FunctionRepo interface {
	FunctionProvider
	Save(ctx context.Context, f *Function) error
}

type FunctionTypeRepo interface {
	FindByID(ctx context.Context, id FunctionTypeID) (FunctionType, error)
	Save(ctx context.Context, functionType FunctionType) error
}

type FunctionEvaluator interface {
	EvaluateFunction(ctx context.Context, sourceCode string, input json.RawMessage) (json.RawMessage, error)
}

func NewFunctionID() FunctionID {
	return FunctionID(uuid.NewString())
}
