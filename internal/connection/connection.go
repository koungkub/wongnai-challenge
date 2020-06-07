package connection

import (
	"database/sql"

	"github.com/pkg/errors"
)

// NewDB create database sql connection to remote server
func NewDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, errors.Wrap(err, "dsn was wrong")
	}

	return db, nil
}
