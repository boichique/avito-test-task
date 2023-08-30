package dbx

import (
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func IsUniqueViolation(err error) bool {
	var perr *pgconn.PgError
	if errors.As(err, &perr) {
		return perr.Code == pgerrcode.UniqueViolation
	}

	return false
}

func IsNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
