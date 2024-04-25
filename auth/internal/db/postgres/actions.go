package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
)

func (p *Postgres) GetUserInfoByID(ctx context.Context, uuid string) (models.UserInfo, error) {
	var user models.UserInfo

	q := `SELECT username FROM users WHERE uuid = $1`

	err := p.QueryRowx(q, uuid).StructScan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return models.UserInfo{}, utils.NewError(utils.ErrNotFound, "user with this uuid not found")
	}
	if err != nil {
		return models.UserInfo{}, utils.NewError(utils.ErrInternal, err.Error())
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

func (p *Postgres) GetUserInfoByEmail(ctx context.Context, email string) (models.UserInfo, error) {
	var user models.UserInfo

	q := `SELECT username, uuid, isverified FROM users WHERE email = $1`

	err := p.QueryRowx(q, email).StructScan(&user)
	if errors.Is(err, sql.ErrNoRows) {
		return models.UserInfo{}, utils.NewError(utils.ErrNotFound, "user with this email not found")
	}
	if err != nil {
		return models.UserInfo{}, utils.NewError(utils.ErrInternal, err.Error())
	}

	return user, nil
}

func (p *Postgres) SaveVerification(ctx context.Context, info models.VerInfo) error {
	q := `INSERT INTO verif (email,code,isverified) VALUES($1,$2,$3)`

	_, err := p.Queryx(q, info.Email, info.Code, info.IsVerified)
	if err != nil {
		return utils.NewError(utils.ErrBadRequest, "user with this email already exists")
	}

	return nil
}

func (p *Postgres) GetVerification(ctx context.Context, email string) (models.VerInfo, error) {
	var info models.VerInfo

	q := `SELECT code,isverified FROM verif WHERE email = $1`

	err := p.QueryRowx(q, email).StructScan(&info)
	if errors.Is(err, sql.ErrNoRows) {
		return models.VerInfo{}, utils.NewError(utils.ErrNotFound, "user with this email not found")
	}
	if err != nil {
		return models.VerInfo{}, utils.NewError(utils.ErrInternal, err.Error())
	}

	return info, nil
}

func (p *Postgres) GetRecover(ctx context.Context, email string) (models.RecoverInfo, error) {
	var info models.RecoverInfo

	q := `SELECT code,password FROM recover WHERE email = $1`

	err := p.QueryRowx(q, email).StructScan(&info)
	if errors.Is(err, sql.ErrNoRows) {
		return models.RecoverInfo{}, utils.NewError(utils.ErrNotFound, "user with this email not found")
	}
	if err != nil {
		return models.RecoverInfo{}, utils.NewError(utils.ErrInternal, err.Error())
	}

	return info, nil
}
