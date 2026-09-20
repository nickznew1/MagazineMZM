package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UpdatePersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	r.logger.Debug("Repository: UpdatePersonalInfo (update personal info for user) started", "input :", input)

	err := r.db.QueryRow(ctx, "UPDATE customer_personal_info SET company =$2, first_name=$3, second_name=$4 WHERE id =$1 RETURNING company, first_name, second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when update personal info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: UpdatePersonalInfo (update personal info for user) success", "input :", input)
	return userInfo, nil
}
