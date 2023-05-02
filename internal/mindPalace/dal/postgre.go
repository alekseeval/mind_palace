package dal

import (
	"MindPalace/internal/mindPalace/configuration"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: daoInterface realization

type PostgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(config *configuration.Config) (*PostgresDB, error) {
	return &PostgresDB{
		db: &sqlx.DB{},
	}, nil
}
