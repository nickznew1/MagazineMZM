package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) RecordDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	r.logger.Debug("Repository: RecordDeliveryInfo started", "input :", input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_delivery_info (id,phone_number, city, address) VALUES ($1, $2, $3, $4) RETURNING id, phone_number, city,address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.Id, &userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when record delivery info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: RecordDeliveryInfo success", "input :", input)
	return userInfo, nil
}
