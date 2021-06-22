package customerrors

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/services/condition"
)

type CodeArchError struct {
	HttpCode int    `json:"-"`
	Message  string `json:"message,omitempty"`
}

func (e *CodeArchError) Error() string {
	ba, _ := json.Marshal(e)
	return string(ba)
}

func NewCodeArchError(code int, message string) error {
	return &CodeArchError{
		HttpCode: code,
		Message:  message,
	}
}

func ErrorResponse(err error) (int, interface{}) {
	e, ok := err.(*CodeArchError)
	if ok {
		return e.HttpCode, e
	}
	return http.StatusInternalServerError, gin.H{"message": err.Error()}
}

type CutomError struct {
	HttpCode   int    `json:"http_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Expression string `json:"expression,omitempty"`
}

type CutomErrors []CutomError

type OverrideErrors []CutomError

func (c *CutomError) Execute(ctx context.Context, m *models.Controller) error {
	cond := condition.Condition{
		Expression: c.Expression,
	}
	status, err := cond.Evalute(ctx, m)
	if err != nil {
		return NewCodeArchError(http.StatusInternalServerError, err.Error())
	}
	if status {
		return NewCodeArchError(c.HttpCode, c.Message)
	}
	return nil
}

func (errs CutomErrors) Execute(ctx context.Context, m *models.Controller) error {
	for _, e := range errs {
		err := e.Execute(ctx, m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (errs OverrideErrors) Execute(ctx context.Context, m *models.Controller, err error) error {
	if err == nil {
		return nil
	}
	for _, e := range errs {
		if strings.Contains(err.Error(), e.Expression) {
			return NewCodeArchError(e.HttpCode, e.Message)
		}
	}
	return err
}
