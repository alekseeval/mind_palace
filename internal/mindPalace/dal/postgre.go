package dal

import (
	"database/sql"
)

// TODO: daoInterface realization

type PostgreDB struct {
	dbConn sql.DB
}

func NewPostgreDB(dbHost string, dbPort int, login string, password string) (*PostgreDB, error) {
	return &PostgreDB{
		dbConn: sql.DB{},
	}, nil
}
