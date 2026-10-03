package users_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfileDeliveryInfo(
	ctx context.Context,
	id string) (model.UserDeliveryInfoOut, error) {

	var profileDeliveryInfo model.UserDeliveryInfoOut

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    SELECT id,phone_number, city, address 
    from customer_delivery_info 
    WHERE id = $1
    `

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&profileDeliveryInfo.Id,
		&profileDeliveryInfo.PhoneNumber,
		&profileDeliveryInfo.City,
		&profileDeliveryInfo.Address,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.UserDeliveryInfoOut{}, nil
		}
		return model.UserDeliveryInfoOut{}, fmt.Errorf("scan query err: %w", err)
	}

	return profileDeliveryInfo, nil
}
