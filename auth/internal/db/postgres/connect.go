package postgres

import (
	"github.com/alserov/hrs/auth/internal/db/migrations"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustConnect(addr string) *sqlx.DB {
	conn, err := sqlx.Connect("postgres", addr)
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	migrations.MustMigrate(conn)

	return conn
}
