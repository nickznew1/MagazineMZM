package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var customer model.UserOrdinaryInfo

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("bcrypt error: %w", err)
	}

	input.Password = string(hashedPassword)

	query := `INSERT INTO customer 
    (login,password, email) VALUES ($1,$2,$3) 
    RETURNING id,login,password,email`

	row := r.pool.QueryRow(ctx, query, input.Login, input.Password, input.Email)

	err = row.Scan(&customer.Id, &customer.Login, &customer.Password, &customer.Email)

	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("scan error: %w", err)
	}

	return customer, nil
}
