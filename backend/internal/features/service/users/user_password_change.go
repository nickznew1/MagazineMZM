package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UserPasswordChange(
	ctx context.Context,
	input model.PasswordChange) (model.PasswordChange, error) {

	user, err := s.UserPasswordChange(ctx, input)
	if err != nil {
		return model.PasswordChange{}, fmt.Errorf("error when trying to change password for user with id='%d': %w", input.Id, err)
	}

	return user, nil
}
