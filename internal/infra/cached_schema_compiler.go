package infra

import (
	"context"
	"crypto/sha256"

	lru "github.com/hashicorp/golang-lru"
	"github.com/jeff-upton-sp/sp-extension/internal/model"
)

type cachedSchemaCompiler struct {
	delegate model.SchemaCompiler
	cache    *lru.Cache
}

func newCachedSchemaCompiler(cacheSize int, delegate model.SchemaCompiler) (*cachedSchemaCompiler, error) {
	cache, err := lru.New(cacheSize)
	if err != nil {
		return nil, err
	}

	c := &cachedSchemaCompiler{}
	c.cache = cache
	c.delegate = delegate

	return c, nil
}

func (c *cachedSchemaCompiler) CompileSchema(ctx context.Context, rawSchema model.RawSchema) (model.Schema, error) {
	sum := sha256.Sum256([]byte(rawSchema))

	if schema, ok := c.cache.Get(sum); ok {
		return schema.(model.Schema), nil
	}

	schema, err := c.delegate.CompileSchema(ctx, rawSchema)
	if err != nil {
		return nil, err
	}

	c.cache.Add(sum, schema)

	return schema, nil
}
