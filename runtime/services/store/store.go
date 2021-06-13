package store

import (
	"context"

	"github.com/Jeffail/gabs/v2"
)

type T int

const (
	PgSQL T = iota
)

type store interface {
	Execute(ctx context.Context, m *gabs.Container) error
}

type Store struct {
	Type       T          `json:"type,omitempty"`
	PgSQLStore PgSQLStore `json:"pg_sql_store,omitempty"`
}

func (l *Store) Execute(ctx context.Context, m *gabs.Container) error {
	err := l.PgSQLStore.Execute(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
