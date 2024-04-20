package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
)

func (p *Postgres) SignUp(ctx context.Context, req models.RegReq) error {
	q := `INSERT INTO users (uuid,email,username,password) VALUES($1,$2,$3,$4)`

	_, err := p.Queryx(q, req.UUID, req.Email, req.UserName, req.Password)
	if err != nil {
		return utils.NewError(utils.ErrBadRequest, "user with this email already exists")
	}
	return nil
}

func (p *Postgres) SignIn(ctx context.Context, email string) (db.LoginInfo, error) {
	var info db.LoginInfo

	q := `SELECT uuid, password, email, username FROM users WHERE email = $1`

	err := p.QueryRowx(q, email).StructScan(&info)

	if errors.Is(err, sql.ErrNoRows) {
		return db.LoginInfo{}, utils.NewError(utils.ErrNotFound, "user with this email not found")
	}
	if err != nil {
		return db.LoginInfo{}, utils.NewError(utils.ErrInternal, err.Error())
	}

	return info, nil
}
