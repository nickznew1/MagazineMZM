package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error) {
	return s.GetAllUsers(ctx)
}
