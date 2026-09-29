package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) RecordPersonalInfo(
	ctx context.Context,
	input model.UserPersonalInfo) (model.UserPersonalInfo, error) {

	var userInfo model.UserPersonalInfo

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    INSERT INTO customer_personal_info
    (id,company, first_name, second_name) 
    VALUES ($1, $2, $3, $4)
    RETURNING id, company, first_name,second_name
    `

	row := r.pool.QueryRow(ctx, query, input.Id, input.Company, input.FirstName, input.SecondName)

	err := row.Scan(
		&userInfo.Id,
		&userInfo.Company,
		&userInfo.FirstName,
		&userInfo.SecondName,
	)

	if err != nil {
		return model.UserPersonalInfo{}, fmt.Errorf("scan query error: %w", err)
	}

	return userInfo, nil
}
