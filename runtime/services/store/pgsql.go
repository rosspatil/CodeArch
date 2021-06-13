package store

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/rosspatil/codearch/runtime/connectors"
	"github.com/rosspatil/codearch/runtime/utils"
)

type PgSQLStore struct {
	Connection string   `json:"connection,omitempty"`
	Query      string   `json:"query,omitempty"`
	Args       []string `json:"args,omitempty"`
}

func (l *PgSQLStore) Execute(ctx context.Context, c *gabs.Container) error {
	db, err := connectors.GetPgSQLConnection(l.Connection)
	if err != nil {
		return err
	}
	args := utils.ResolveValues(l.Args, c)
	_, err = db.ExecContext(ctx, l.Query, args...)
	return err
}
