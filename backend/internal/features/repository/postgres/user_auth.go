package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) UserAuth(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	r.logger.Debug("Repository: UserAuth started", "user_id: ", input.Id)

	user, err := r.GetUserById(ctx, input)
	if err != nil {
		r.logger.Error("Repository: UserAuth error - can t get user info when check user_id", slog.Any("db_err: ", err))
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		r.logger.Error("Repository: UserAuth error - wrong password for user", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: UserAuth success", "user_id: ", input.Id)
	return user, nil
}
