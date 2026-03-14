package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type functionTypeRepo struct {
	mu    sync.RWMutex
	types map[model.FunctionTypeID]model.FunctionType
}

func NewFunctionTypeRepo() *functionTypeRepo {
	r := &functionTypeRepo{}
	r.types = make(map[model.FunctionTypeID]model.FunctionType)

	return r
}

func (r *functionTypeRepo) FindByID(ctx context.Context, id model.FunctionTypeID) (model.FunctionType, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	functionType, ok := r.types[id]
	if !ok {
		return model.FunctionType{}, fmt.Errorf("function type '%s' not found", id)
	}

	return functionType, nil
}

func (r *functionTypeRepo) Save(ctx context.Context, functionType model.FunctionType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.types[functionType.ID] = functionType

	return nil
}
