package load

import (
	"context"

	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/services/customerrors"
)

type T int

const (
	PgSQL T = iota
)

type load interface {
	Execute(ctx context.Context, c *models.Controller) ([]interface{}, error)
}

type Load struct {
	Type        T                        `json:"type,omitempty"`
	ResultField string                   `json:"result_field,omitempty"`
	PgSQLLoad   PgSQLLoad                `json:"pg_sql_load,omitempty"`
	CutomErrors customerrors.CutomErrors `json:"cutom_errors,omitempty"`
}

func (l *Load) Execute(ctx context.Context, m *models.Controller) error {
	data, err := l.PgSQLLoad.Execute(ctx, m)
	if err != nil {
		return err
	}
	m.SetP(data, l.ResultField)
	return nil
}
