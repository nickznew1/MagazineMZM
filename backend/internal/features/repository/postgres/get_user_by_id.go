package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) GetUserById(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	var user model.UserOrdinaryInfo

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
     SELECT id,login, password 
     FROM customer 
     WHERE login =$1
    `

	row := r.pool.QueryRow(ctx, query, input.Login)

	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Password,
	)

	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("scan query err: %w", err)
	}

	return user, nil
}
