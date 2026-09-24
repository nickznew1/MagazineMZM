package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	user, err := s.CreateUser(ctx, input)
	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("error when trying to create user(service): %w", err)
	}
	return user, nil
}
