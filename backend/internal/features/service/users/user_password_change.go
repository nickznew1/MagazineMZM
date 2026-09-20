package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error) {
	return s.UserPasswordChange(ctx, input)
}
