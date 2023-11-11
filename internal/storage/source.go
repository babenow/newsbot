package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SourcePostgresStorage struct {
	db *sqlx.DB
}
