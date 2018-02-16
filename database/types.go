package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	// this is a comment justifying the very standard way of importing a db driver
	_ "github.com/lib/pq"
)

type myDB interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	QueryRowx(query string, args ...interface{}) *sqlx.Row
}

// Client struct wraps sqlx.db so we can write functions like client.LoadUsers()
type Client struct {
	myDB
}

// New function
func New(user, pass, host, port, db string) Client {
	connStrTemplate := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	connStr := fmt.Sprintf(connStrTemplate, user, pass, host, port, db)

	return Client{sqlx.MustConnect("postgres", connStr)}
}
