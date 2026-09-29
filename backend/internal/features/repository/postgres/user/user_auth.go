package users_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRepository) UserAuth(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	user, err := r.GetUserById(ctx, input)
	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("repo error:%w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("bcrypt error: %w", err)
	}

	return user, nil
}
