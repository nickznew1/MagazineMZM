package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	r.logger.Debug("Repository: UpdateDeliveryInfo started", "input :", input)
	err := r.db.QueryRow(ctx, "UPDATE customer_delivery_info SET phone_number=$2, city=$3, address=$4 WHERE id =$1 RETURNING phone_number, city, address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when update delivery info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: UpdateDeliveryInfo success", "input :", input)
	return userInfo, nil
}
