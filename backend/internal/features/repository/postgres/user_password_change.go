package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error) {
	var check model.PasswordChange
	r.logger.Debug("Repository: UserPasswordChange started", "input :", input)
	err := r.db.QueryRow(ctx, "SELECT password FROM customer WHERE id =$1", input.Id).Scan(&check.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(check.OldPassword), []byte(input.OldPassword))
	if err != nil {
		r.logger.Error("Repository: UserPasswordChange bcrypt error - user writes wrong ordinary password", slog.Any("bcrypt_err: ", err))
		return check, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 12)
	err = r.db.QueryRow(ctx, "UPDATE customer SET password = $1 WHERE id = $2 RETURNING password,id", hashedPassword, input.Id).Scan(&check.NewPassword, &check.Id)
	if err != nil {
		r.logger.Error("Repository: UserPasswordChange bcrypt error when update password for user", slog.Any("db_err: ", err))
		return check, err
	}
	r.logger.Debug("Repository: UserPasswordChange success", "input :", input)
	return check, nil
}
