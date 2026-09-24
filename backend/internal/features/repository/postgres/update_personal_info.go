package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UpdatePersonalInfo(
	ctx context.Context,
	input model.UserPersonalInfo) (model.UserPersonalInfo, error) {

	var userInfo model.UserPersonalInfo

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query :=
		`
    UPDATE customer_personal_info 
    SET company =$2, first_name=$3, second_name=$4 
    WHERE id =$1 
    RETURNING company, first_name, second_name
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
