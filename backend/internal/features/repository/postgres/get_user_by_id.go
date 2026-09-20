package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	r.logger.Debug("Repository: GetUserById (Getting user by login) started", "user_login: ", input.Login)
	err := r.db.QueryRow(ctx, "SELECT id,login, password FROM customer WHERE login =$1", input.Login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		r.logger.Error("Repository: GetUserById (Getting user by login) error", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: GetUserById (Getting user by login) success", "user_login: ", input.Login)
	return user, nil
}
