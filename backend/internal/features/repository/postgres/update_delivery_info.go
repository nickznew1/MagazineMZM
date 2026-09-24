package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UpdateDeliveryInfo(
	ctx context.Context,
	input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var userInfo model.UserDeliveryInfo

	query := `
    UPDATE customer_delivery_info 
    SET phone_number=$2, city=$3, address=$4 
    WHERE id =$1 
    RETURNING phone_number, city, address
    `

	row := r.pool.QueryRow(ctx, query, input.Id, input.PhoneNumber, input.City, input.Address)

	err := row.Scan(
		&userInfo.PhoneNumber,
		&userInfo.City,
		&userInfo.Address,
	)

	if err != nil {
		return model.UserDeliveryInfo{}, fmt.Errorf("scan query error: %w", err)
	}

	return userInfo, nil
}
