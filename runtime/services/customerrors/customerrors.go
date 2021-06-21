package customerrors

import (
	"context"
	"net/http"

	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/services/condition"
)

type CutomError struct {
	HttpCode   int
	Message    string
	Expression string
}

type CutomErrors []CutomError

func (c *CutomError) Execute(ctx context.Context, m *models.Controller) (int, interface{}) {
	cond := condition.Condition{}
	status, err := cond.Evalute(ctx, m)
	if err != nil {
		return http.StatusInternalServerError, map[string]string{"message": err.Error()}
	}
	if status {
		return c.HttpCode, map[string]string{"message": c.Message}
	}
	return 0, nil
}

func (errs CutomErrors) Execute(ctx context.Context, m *models.Controller) (int, interface{}) {
	for _, e := range errs {
		code, err := e.Execute(ctx, m)
		if err != nil {
			return code, err
		}
	}
	return 0, nil
}
