package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error) {
	var profileInfo model.UserOrdinaryInfoOut
	r.logger.Debug("Repository: FetchProfileInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,login,email,user_role from customer WHERE id =$1", id).Scan(&profileInfo.Id, &profileInfo.Login, &profileInfo.Email, &profileInfo.UserRole)
	if err != nil {
		r.logger.Error("Repository: FetchProfileInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profileInfo, err
	}
	r.logger.Debug("Repository: FetchProfileInfo success (goroutines) ", "user_id: ", id)
	return profileInfo, nil
}
