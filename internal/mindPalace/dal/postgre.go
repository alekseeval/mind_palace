package dal

import (
	"MindPalace/internal/mindPalace/configuration"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: daoInterface realization

type PostgresDB struct {
	db *sqlx.DB
}

// TODO: Additional dbconn settings (maxconns, timeout)
func NewPostgresDB(config *configuration.Config) (*PostgresDB, error) {
	dbConfig := config.System.DB
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	return &PostgresDB{
		db: db,
	}, err
}
