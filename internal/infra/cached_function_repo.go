package infra

import (
	"context"
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type cachedFunctionRepo struct {
	cache    *expirable.LRU[model.FunctionID, model.Function]
	delegate model.FunctionRepo
}

func newCachedFunctionRepo(cacheSize int, cacheDuration time.Duration, delegate model.FunctionRepo) (*cachedFunctionRepo, error) {
	r := &cachedFunctionRepo{}
	r.cache = expirable.NewLRU[model.FunctionID, model.Function](cacheSize, nil, cacheDuration)
	r.delegate = delegate

	return r, nil
}

func (r *cachedFunctionRepo) FindByID(ctx context.Context, id model.FunctionID) (model.Function, error) {
	if f, ok := r.cache.Get(id); ok {
		return f, nil
	}

	f, err := r.delegate.FindByID(ctx, id)
	if err != nil {
		return model.Function{}, err
	}

	r.cache.Add(id, f)

	return f, nil
}

func (r *cachedFunctionRepo) Save(ctx context.Context, f *model.Function) error {
	return r.delegate.Save(ctx, f)
}
