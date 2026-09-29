package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) GetAllUsers(
	ctx context.Context) ([]model.UserOrdinaryInfo, error) {

	var users []model.UserOrdinaryInfo

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    SELECT id,login,password,email, registration_date,user_role 
    FROM customer
    `
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return []model.UserOrdinaryInfo{}, fmt.Errorf("query rows error: %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var user model.UserOrdinaryInfo
		err = rows.Scan(
			&user.Id,
			&user.Login,
			&user.Password,
			&user.Email,
			&user.RegistrationDate,
			&user.UserRole)
		if err != nil {
			return []model.UserOrdinaryInfo{}, fmt.Errorf("scan query error: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
