package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/alserov/hrs/auth/internal/db"
	"github.com/alserov/hrs/auth/internal/log"
	"github.com/alserov/hrs/auth/internal/models"
	"github.com/alserov/hrs/auth/internal/utils"
)

func (p *Postgres) SignUp(ctx context.Context, req models.RegReq) error {
	q := `INSERT INTO users (uuid,email,username,password,isverified) VALUES($1,$2,$3,$4,$5)`

	_, err := p.Queryx(q, req.UUID, req.Email, req.UserName, req.Password, false)
	if err != nil {
		return utils.NewError(utils.ErrBadRequest, "user with this email already exists")
	}

	log.GetLogger(ctx).Debug("repo layer success ✔")

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

	log.GetLogger(ctx).Debug("repo layer success ✔")

	return info, nil
}

func (p *Postgres) ResetPass(ctx context.Context, uuid string, password string) error {
	q := `UPDATE users SET password = $1 WHERE uuid = $2`

	_, err := p.Queryx(q, password, uuid)
	if err != nil {
		return utils.NewError(utils.ErrNotFound, "user with this id not found")
	}

	log.GetLogger(ctx).Debug("repo layer success ✔")

	return nil
}

func (p *Postgres) RecoverPass(ctx context.Context, req models.RecoverReq) error {
	q := `INSERT INTO recover (email,password,code) VALUES($1,$2,$3)`

	_, err := p.Queryx(q, req.Email, req.NewPassword, req.Code)

	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo layer success ✔")

	return nil
}

func (p *Postgres) Verificate(ctx context.Context, code, email string) error {
	q1 := `DELETE FROM verif WHERE email = $1 AND code = $2`

	_, err := p.Queryx(q1, email, code)
	if err != nil {
		return utils.NewError(utils.ErrNotFound, "wrong code or email provided")
	}

	q2 := `UPDATE users SET isverified = $1 WHERE email = $2`

	_, err = p.Queryx(q2, true, email)
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo layer success ✔")

	return nil
}

func (p *Postgres) VerificateRecover(ctx context.Context, code, email, password string) error {
	q1 := `DELETE FROM recover WHERE email = $1 AND code = $2`

	_, err := p.Queryx(q1, email, code)
	if err != nil {
		return utils.NewError(utils.ErrNotFound, "wrong code or email provided")
	}

	q2 := `UPDATE users SET password = $1 WHERE email = $2`

	_, err = p.Queryx(q2, password, email)
	if err != nil {
		return utils.NewError(utils.ErrInternal, err.Error())
	}

	log.GetLogger(ctx).Debug("repo layer success ✔")

	return nil
}
