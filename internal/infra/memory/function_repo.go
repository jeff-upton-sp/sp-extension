package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type functionRepo struct {
	mu    sync.RWMutex
	store map[model.FunctionID]model.Function
}

func NewFunctionRepo() (*functionRepo, error) {
	r := &functionRepo{}
	r.store = make(map[model.FunctionID]model.Function)

	return r, nil
}

func (r *functionRepo) FindByID(ctx context.Context, id model.FunctionID) (model.Function, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if f, ok := r.store[id]; ok {
		return f, nil
	}

	return model.Function{}, fmt.Errorf("function '%s' not found", id)
}

func (r *functionRepo) Save(ctx context.Context, f *model.Function) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if f.ID == "" {
		f.ID = model.NewFunctionID()
	}

	r.store[f.ID] = *f

	return nil
}
