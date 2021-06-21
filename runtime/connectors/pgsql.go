package connectors

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rosspatil/codearch/runtime/utils"
)

type pgsqlConnectors map[string]*sql.DB

var (
	pgsqlconnections pgsqlConnectors
)

func init() {
	pgsqlconnections = make(pgsqlConnectors)
}

type PGSQL struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func (m *PGSQL) ConnectionString() string {
	m.User, _ = utils.ResolveEnvironmentVariable(m.User)
	m.Password, _ = utils.ResolveEnvironmentVariable(m.Password)
	m.DB, _ = utils.ResolveEnvironmentVariable(m.DB)
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", m.User, m.Password, m.DB)
}

func InitPgSQL(ctx context.Context, m *PGSQL) error {
	fmt.Println("here in pgsql connector")
	db, err := sql.Open("postgres", m.ConnectionString())
	if err != nil {
		return err
	}
	err = db.PingContext(ctx)
	if err != nil {
		return err
	}
	pgsqlconnections[m.Name] = db
	return nil
}

func GetPgSQLConnection(name string) (*sql.DB, error) {
	db, ok := pgsqlconnections[name]
	if !ok {
		return nil, fmt.Errorf("connection with name '%s' not found", name)
	}
	return db, nil
}
