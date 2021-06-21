package condition

import (
	"context"
	"fmt"
	"reflect"

	"github.com/PaesslerAG/gval"
	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/utils"
)

type Condition struct {
	Expression  string            `json:"query,omitempty"`
	Args        map[string]string `json:"args,omitempty"`
	ResultField string            `json:"result_field,omitempty"`
}

func (l *Condition) Execute(ctx context.Context, c *models.Controller) error {
	status, err := l.Evalute(ctx, c)
	if err != nil {
		return err
	}
	c.SetP(status, l.ResultField)
	return nil
}

func (l *Condition) Evalute(ctx context.Context, c *models.Controller) (bool, error) {
	f := gval.Full(ArrarLenLanguage(), resolve(c))
	evaluable, err := f.NewEvaluable(l.Expression)
	if err != nil {
		return false, err
	}
	m := map[string]interface{}{}
	for k, v := range l.Args {
		values := utils.ResolveValue(v, c)
		m[k] = values
	}
	v, err := evaluable(ctx, m)
	if err != nil {
		return false, err
	}
	status, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("expressoins output is not boolean")
	}
	return status, nil
}

func ArrarLenLanguage() gval.Language {
	return gval.NewLanguage(gval.JSON(), gval.Arithmetic(), gval.Function("arrayLen", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return fmt.Errorf("arrayLen function expect exactly one argument"), nil
		}
		if arguments[0] == nil {
			return 0, nil
		}
		t := reflect.TypeOf(arguments[0])
		if t.Kind() != reflect.Slice {
			return fmt.Errorf("arrayLen function expect array"), nil
		}
		return reflect.ValueOf(arguments[0]).Len(), nil
	}))
}

func resolve(c *models.Controller) gval.Language {
	return gval.NewLanguage(gval.JSON(), gval.Function("resolve", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return fmt.Errorf("resolve function expect exactly one argument"), nil
		}
		arg, ok := arguments[0].(string)
		if !ok {
			return fmt.Errorf("resolve function expect only string type argument"), nil
		}
		return c.Path(arg).Data(), nil
	}))
}
