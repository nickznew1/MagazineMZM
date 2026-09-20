package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) RecordPersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	r.logger.Debug("Repository: RecordPersonalInfo (record new personal info for user) started", "input :", input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_personal_info (id,company, first_name, second_name) VALUES ($1, $2, $3, $4) RETURNING id, company, first_name,second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Id, &userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when insert new personal info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: RecordPersonalInfo (record new personal info for user) success", "input :", input)
	return userInfo, nil
}
