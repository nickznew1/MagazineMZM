package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	var profilePersonalInfo model.UserPersonalInfoOut
	r.logger.Debug("Repository: FetchProfilePersonalInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,company, first_name, second_name from customer_personal_info WHERE id=$1", id).Scan(&profilePersonalInfo.Id, &profilePersonalInfo.Company, &profilePersonalInfo.FirstName, &profilePersonalInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: FetchProfilePersonalInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profilePersonalInfo, err
	}
	r.logger.Debug("Repository: FetchProfilePersonalInfo success (goroutines)", "user_id: ", id)
	return profilePersonalInfo, nil
}
