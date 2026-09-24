package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UserChangeEmail(
	ctx context.Context,
	input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {

	user, err := s.UserChangeEmail(ctx, input)
	if err != nil {
		return model.UserOrdinaryInfo{}, fmt.Errorf("error when trying to change email for user with id='%d': %w", input.Id, err)
	}

	return user, nil
}
