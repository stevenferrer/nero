package nero

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// SQLRunner is an interface that wraps the standard sql methods
type SQLRunner interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// ValueScanner is an interface that wraps the driver.Valuer and sql.Scanner interface
type ValueScanner interface {
	driver.Valuer
	sql.Scanner
}
