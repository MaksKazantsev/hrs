package postgres

import (
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	*sqlx.DB
}

var _ db.Repository = &Postgres{}

func NewRepository(db *sqlx.DB) *Postgres {
	return &Postgres{
		db,
	}
}
