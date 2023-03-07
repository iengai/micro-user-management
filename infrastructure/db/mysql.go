package db

import (
	"database/sql"
)

func NewMySQLDB(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}
