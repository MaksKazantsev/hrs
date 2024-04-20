package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
)

func (p *Postgres) GetUserInfo(ctx context.Context, uuid string) (models.User, error) {
	var user models.User

	q := `SELECT email, password, username FROM users WHERE uuid = $1`

	err := p.QueryRowx(q, uuid).StructScan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return models.User{}, utils.NewError(utils.ErrNotFound, "user with this uuid not found")
	}
	if err != nil {
		return models.User{}, utils.NewError(utils.ErrInternal, err.Error())
	}

	return user, nil
}

func (p *Postgres) GetUserPassword(ctx context.Context, uuid string) (string, error) {
	var password string
	q := `SELECT password FROM users WHERE uuid = $1`

	err := p.QueryRowx(q, uuid).Scan(&password)
	if errors.Is(err, sql.ErrNoRows) {
		return "", utils.NewError(utils.ErrNotFound, "user with this uuid not found")
	}
	if err != nil {
		return "", utils.NewError(utils.ErrInternal, err.Error())
	}

	return password, nil

}
