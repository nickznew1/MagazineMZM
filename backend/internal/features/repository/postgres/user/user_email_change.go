package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) UserChangeEmail(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var user model.UserOrdinaryInfo
	var loginExists bool

	query := `
     SELECT EXISTS 
    (SELECT 1 FROM customer 
    WHERE login =$1)
   `

	row := r.pool.QueryRow(ctx, query, input.Login)

	err := row.Scan(
		&loginExists,
	)

	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("scan query error:%w", err)
	}

	if !loginExists {
		return model.UserOrdinaryInfo{}, fmt.Errorf("user with id='%d' not exists", input.Id)
	}

	query = `
    UPDATE customer 
    SET email = $1 
    WHERE login = $2 
    RETURNING login,email
     `

	row = r.pool.QueryRow(ctx, query, input.Email, input.Login)

	err = row.Scan(
		&user.Login,
		&user.Email,
	)

	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("scan query error: %w", err)
	}
	return user, nil
}
