package infra

import (
	"context"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type jsonSchema struct {
	schema *gojsonschema.Schema
}

func newJSONSchema(jsonBytes []byte) (*jsonSchema, error) {
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("parse schema: %w", err)
	}

	s := &jsonSchema{}
	s.schema = schema

	return s, nil
}

func (s *jsonSchema) ValidateJSON(ctx context.Context, data []byte) error {
	result, err := s.schema.Validate(gojsonschema.NewBytesLoader(data))
	if err != nil {
		return fmt.Errorf("validate jsona: %w", err)
	}

	var errSB strings.Builder

	if !result.Valid() {
		for _, err := range result.Errors() {
			errSB.WriteString(err.String())
			errSB.WriteString("\n")
		}

		return fmt.Errorf("validation failed: %s", errSB.String())
	}

	return nil
}
