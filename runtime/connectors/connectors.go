package connectors

import (
	"context"
)

type T int

const (
	PgSQLConnector T = iota
)

type Connector struct {
	Type           T     `json:"type,omitempty"`
	PgSQLConnector PGSQL `json:"data,omitempty"`
}

func (c *Connector) Load(ctx context.Context) error {
	switch c.Type {
	case PgSQLConnector:
		return InitPgSQL(ctx, &c.PgSQLConnector)
	}
	return nil
}
