package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) RecordDeliveryInfo(
	ctx context.Context,
	input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {

	var userInfo model.UserDeliveryInfo

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
     INSERT INTO customer_delivery_info 
     (id,phone_number, city, address) 
     VALUES ($1, $2, $3, $4) 
     RETURNING id, phone_number, city,address
    `

	row := r.pool.QueryRow(ctx, query, input.Id, input.PhoneNumber, input.City, input.Address)

	err := row.Scan(
		&userInfo.Id,
		&userInfo.PhoneNumber,
		&userInfo.City,
		&userInfo.Address,
	)

	if err != nil {
		return model.UserDeliveryInfo{}, fmt.Errorf("scan query error: %w", err)
	}

	return userInfo, nil
}
