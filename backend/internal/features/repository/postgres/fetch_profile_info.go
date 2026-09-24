package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) FetchProfileInfo(
	ctx context.Context,
	id string) (model.UserOrdinaryInfoOut, error) {

	var profileInfo model.UserOrdinaryInfoOut

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
     SELECT id,login,email,user_role 
     from customer 
     WHERE id =$1
    `
	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&profileInfo.Id,
		&profileInfo.Login,
		&profileInfo.Email,
		&profileInfo.UserRole,
	)
	if err != nil {
		return model.UserOrdinaryInfoOut{}, fmt.Errorf("scan query err: %w", err)
	}

	return profileInfo, nil
}
