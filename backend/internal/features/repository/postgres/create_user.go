package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) CreateUser(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var customer model.UserOrdinaryInfo
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	r.logger.Debug("Repository: CreateUser started")
	if err != nil {
		r.logger.Error("Repository: CreateUser error - can't hash password(bcrypt)", slog.Any("bcrypt_err: ", err))
		return customer, err
	}
	input.Password = string(hashedPassword)

	err = r.db.QueryRow(ctx,
		"INSERT INTO customer (login,password, email) VALUES ($1,$2,$3) RETURNING id,login,password,email",
		input.Login, input.Password, input.Email).
		Scan(&customer.Id, &customer.Login, &customer.Password, &customer.Email)
	if err != nil {
		r.logger.Error("Repository: CreateUser error when insert new user into customer table", slog.Any("db_err: ", err))
		return customer, err
	}

	r.logger.Debug("Repository: CreateUser success", "new user_id: ", customer.Id)
	return customer, nil
}
