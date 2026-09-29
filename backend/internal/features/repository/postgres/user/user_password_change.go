package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) UserPasswordChange(
	ctx context.Context,
	input model.PasswordChange) (model.PasswordChange, error) {
	var check model.PasswordChange
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    SELECT password 
    FROM customer 
    WHERE id =$1
    `

	row := r.pool.QueryRow(ctx, query, input.Id)

	err := row.Scan(
		&check.OldPassword,
	)

	if err != nil {
		return model.PasswordChange{}, fmt.Errorf("scan query error:%w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(check.OldPassword), []byte(input.OldPassword))
	if err != nil {
		return model.PasswordChange{}, fmt.Errorf("bcrypt compare error: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 12)
	if err != nil {
		return model.PasswordChange{}, fmt.Errorf("bcrypt generate error: %w", err)
	}

	query = `
    UPDATE customer 
    SET password = $1 
    WHERE id = $2 
    RETURNING password,id
    `

	row = r.pool.QueryRow(ctx, query, hashedPassword, input.Id)

	err = row.Scan(
		&check.NewPassword,
		&check.Id,
	)

	if err != nil {
		return model.PasswordChange{}, fmt.Errorf("scan query error: %w", err)
	}

	return check, nil
}
