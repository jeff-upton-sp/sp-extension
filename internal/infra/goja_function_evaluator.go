package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

	// TODO: is there safe way to cache VMs? keeping them warm?
	// TODO:  memory limit..
	vm := goja.New()

	// Execution time limit enforced...
	time.AfterFunc(300*time.Millisecond, func() {
		vm.Interrupt("halt")
	})

	_, err := vm.RunString(sourceCode)
	if err != nil {
		return nil, err
	}

	main, ok := goja.AssertFunction(vm.Get("main"))
	if !ok {
		return nil, fmt.Errorf("no main function defined")
	}

	result, err := main(goja.Undefined(), vm.ToValue(inputMap))
	if err != nil {
		return nil, fmt.Errorf("eval main: %w", err)
	}

	js, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("marshal result: %w", err)
	}

	return json.RawMessage(js), nil
}
