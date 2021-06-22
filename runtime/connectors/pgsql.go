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
	Name     string `json:"name,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DB       string `json:"db,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

func (m *PGSQL) ConnectionString() string {
	m.User, _ = utils.ResolveEnvironmentVariable(m.User)
	m.Password, _ = utils.ResolveEnvironmentVariable(m.Password)
	m.DB, _ = utils.ResolveEnvironmentVariable(m.DB)
	m.Host, _ = utils.ResolveEnvironmentVariable(m.Host)
	m.Port, _ = utils.ResolveEnvironmentVariable(m.Port)
	fmt.Println(fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s  sslmode=disable", m.User, m.Password, m.DB, m.Host, m.Port))
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s  sslmode=disable", m.User, m.Password, m.DB, m.Host, m.Port)
}

func InitPgSQL(ctx context.Context, m *PGSQL) error {
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
