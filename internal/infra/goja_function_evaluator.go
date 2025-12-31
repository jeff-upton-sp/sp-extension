package infra

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dop251/goja"
)

type gojaFunctionEvaluator struct {
}

func newGojaFunctionEvaluator() *gojaFunctionEvaluator {
	e := &gojaFunctionEvaluator{}
	return e
}

func (e *gojaFunctionEvaluator) EvaluateFunction(ctx context.Context, sourceCode string, input json.RawMessage) (json.RawMessage, error) {
	var inputMap map[string]any
	if err := json.Unmarshal(input, &inputMap); err != nil {
		return nil, fmt.Errorf("marshal input: %w", err)
	}

	// TODO: execution time limit, memory limit..

	vm := goja.New()
	vm.Set("input", inputMap)

	result, err := vm.RunString(sourceCode)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("marshal result: %w", err)
	}

	return json.RawMessage(js), nil
}
