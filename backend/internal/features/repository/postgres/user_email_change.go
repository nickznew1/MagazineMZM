package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UserChangeEmail(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	var loginExists bool
	r.logger.Debug("Repository: UserChangeEmail started", "input:", input)
	err := r.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM customer WHERE login =$1)", input.Login).Scan(&loginExists)
	if err != nil {
		r.logger.Error("Repository: UserChangeEmail error when update email for user - login doesn't found", slog.Any("db_err: ", err))
		return user, err
	}
	if !loginExists {
		r.logger.Error("Repository: UserChangeEmail error when update email for user - login doesn't found", slog.Any("db_err: ", err))
		return user, err
	}
	err = r.db.QueryRow(ctx, "UPDATE customer SET email = $1 WHERE login = $2 RETURNING login,email", input.Email, input.Login).Scan(&user.Login, &user.Email)
	if err != nil {
		r.logger.Error("Repository: UserChangeEmail error when update email for user (db error)", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: UserChangeEmail success", "input:", input)
	return user, nil
}
