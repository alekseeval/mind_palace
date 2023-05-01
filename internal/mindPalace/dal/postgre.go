package dal

import (
	"github.com/jmoiron/sqlx"
)

// TODO: daoInterface realization

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) (*PostgresDB, error) {
	return &PostgresDB{
		db: db,
	}, nil
}
