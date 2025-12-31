package memory

import (
	"context"
	"fmt"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type functionRepo struct {
	store map[model.FunctionID]model.Function
}

func NewFunctionRepo() (*functionRepo, error) {
	r := &functionRepo{}
	r.store = make(map[model.FunctionID]model.Function)

	return r, nil
}

func (r *functionRepo) FindByID(ctx context.Context, id model.FunctionID) (model.Function, error) {
	if f, ok := r.store[id]; ok {
		return f, nil
	}

	return model.Function{}, fmt.Errorf("function '%s' not found", id)
}
