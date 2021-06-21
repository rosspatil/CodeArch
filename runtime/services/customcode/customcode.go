package customcode

import (
	"context"
	"fmt"

	"github.com/rosspatil/code-arch/types"
	"github.com/rosspatil/codearch/runtime/models"
)

type CustomCode struct {
	FunctionName string `json:"function_name,omitempty"`
	Data         string `json:"data,omitempty"`
}

func Compiler() {
}

var customCodeRegistry = map[string]CodeFn{}

func RegisterCustomCode(name string, fn CodeFn) {
	customCodeRegistry[name] = fn
}

type CodeFn func(ctx context.Context, c types.Controller) error

func (l *CustomCode) Execute(ctx context.Context, c *models.Controller) error {
	fn, ok := customCodeRegistry[l.FunctionName]
	if !ok {
		return fmt.Errorf("no CustomCode found for %s", l.FunctionName)
	}
	return fn(ctx, c)
}
