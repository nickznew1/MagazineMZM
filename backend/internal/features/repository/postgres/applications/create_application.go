package applications_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ApplicationRepository) CreateApplication(
	ctx context.Context,
	input model.Application) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var id string
	date := time.Now()
	items, err := json.Marshal(input.Items)
	if err != nil {
		return "", fmt.Errorf("json marshal error :%w", err)
	}

	query := `
    INSERT INTO 
    user_applications(user_id,
                      email, 
                      first_name, 
                      second_name, 
                      login, 
                      phone_number, 
                      company, 
                      address,
                      city, 
                      order_date, 
                      items) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    RETURNING id
    `

	row := r.pool.QueryRow(
		ctx,
		query,
		input.UserId,
		input.Email,
		input.FirstName,
		input.SecondName,
		input.Login,
		input.PhoneNumber,
		input.Company,
		input.Address,
		input.City,
		date,
		items,
	)

	err = row.Scan(
		&id,
	)

	if err != nil {
		return "", fmt.Errorf("scan query err : %w", err)
	}

	return id, nil
}
