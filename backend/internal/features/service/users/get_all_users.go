package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) GetAllUsers(
	ctx context.Context) (
	[]model.UserOrdinaryInfo, error,
) {
	user, err := s.GetAllUsers(ctx)
	if err != nil {
		return []model.UserOrdinaryInfo{}, fmt.Errorf("error when getting all users: %w", err)
	}
	return user, nil
}
