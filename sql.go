package nero

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// SQLRunner is the standard sql interface runner
type SQLRunner interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Exec(string, ...interface{}) (sql.Result, error)
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

// ValueScanner is the composition of driver.Valuer and sql.Scanner interfaces
type ValueScanner interface {
	driver.Valuer
	sql.Scanner
}
