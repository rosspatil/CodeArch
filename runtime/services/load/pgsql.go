package load

import (
	"context"

	"github.com/rosspatil/codearch/runtime/connectors"
	"github.com/rosspatil/codearch/runtime/models"
	"github.com/rosspatil/codearch/runtime/utils"
)

type PgSQLLoad struct {
	Connection string   `json:"connection,omitempty"`
	Query      string   `json:"query,omitempty"`
	Args       []string `json:"args,omitempty"`
}

func (l *PgSQLLoad) Execute(ctx context.Context, c *models.Controller) ([]interface{}, error) {
	rs := []interface{}{}
	db, err := connectors.GetPgSQLConnection(l.Connection)
	if err != nil {
		return rs, err
	}
	args := utils.ResolveValues(l.Args, c)
	rows, err := db.QueryContext(ctx, l.Query, args...)
	if err != nil {
		return rs, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return rs, err
	}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return rs, err
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		rs = append(rs, m)
	}
	return rs, nil
}
