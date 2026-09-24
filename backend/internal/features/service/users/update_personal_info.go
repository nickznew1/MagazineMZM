package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UpdatePersonalInfo(
	ctx context.Context,
	input model.UserPersonalInfo) (model.UserPersonalInfo, error) {

	user, err := s.UpdatePersonalInfo(ctx, input)
	if err != nil {
		return model.UserPersonalInfo{}, fmt.Errorf("error when update personal info for id='%d': %w", input.Id, err)
	}
	return user, nil
}
