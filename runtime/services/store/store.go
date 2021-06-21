package store

import (
	"context"

	"github.com/rosspatil/codearch/runtime/models"
)

type T int

const (
	PgSQL T = iota
)

type store interface {
	Execute(ctx context.Context, m *models.Controller) error
}

type Store struct {
	Type       T          `json:"type,omitempty"`
	PgSQLStore PgSQLStore `json:"pg_sql_store,omitempty"`
}

func (l *Store) Execute(ctx context.Context, m *models.Controller) error {
	err := l.PgSQLStore.Execute(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
