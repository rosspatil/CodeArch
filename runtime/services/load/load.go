package load

import (
	"context"

	"github.com/Jeffail/gabs/v2"
)

type T int

const (
	PgSQL T = iota
)

type load interface {
	Execute(ctx context.Context, c *gabs.Container) ([]interface{}, error)
}

type Load struct {
	Type        T         `json:"type,omitempty"`
	ResultField string    `json:"result_field,omitempty"`
	PgSQLLoad   PgSQLLoad `json:"pg_sql_load,omitempty"`
}

func (l *Load) Execute(ctx context.Context, m *gabs.Container) error {
	data, err := l.PgSQLLoad.Execute(ctx, m)
	if err != nil {
		return err
	}
	m.SetP(data, l.ResultField)
	return nil
}
