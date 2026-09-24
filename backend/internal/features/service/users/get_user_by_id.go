package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) GetUserById(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	user, err := s.GetUserById(ctx, input)
	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("error when trying to get user with id='%d': %w", input.Id, err)
	}
	return user, nil
}
