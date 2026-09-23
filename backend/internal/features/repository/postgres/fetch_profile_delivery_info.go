package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error) {
	var profileDeliveryInfo model.UserDeliveryInfoOut
	r.logger.Debug("Repository: FetchProfileDeliveryInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,phone_number, city, address from customer_delivery_info WHERE id = $1", id).Scan(&profileDeliveryInfo.Id, &profileDeliveryInfo.PhoneNumber, &profileDeliveryInfo.City, &profileDeliveryInfo.Address)
	if err != nil {
		r.logger.Error("Repository: FetchProfileDeliveryInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profileDeliveryInfo, err
	}
	r.logger.Debug("Repository: FetchProfileDeliveryInfo success (goroutines)", "user_id: ", id)
	return profileDeliveryInfo, nil
}
