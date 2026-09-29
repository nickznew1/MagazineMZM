package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfilePersonalInfo(
	ctx context.Context,
	id string) (model.UserPersonalInfoOut, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var profilePersonalInfo model.UserPersonalInfoOut

	query := `
    SELECT id,company, first_name, second_name 
    from customer_personal_info 
    WHERE id=$1`

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&profilePersonalInfo.Id,
		&profilePersonalInfo.Company,
		&profilePersonalInfo.FirstName,
		&profilePersonalInfo.SecondName,
	)
	if err != nil {
		return model.UserPersonalInfoOut{}, fmt.Errorf("scan query error: %w", err)
	}

	return profilePersonalInfo, nil
}
