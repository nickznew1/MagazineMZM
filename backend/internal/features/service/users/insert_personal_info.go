package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) RecordPersonalInfo(
	ctx context.Context,
	input model.UserPersonalInfo) (model.UserPersonalInfo, error) {

	user, err := s.RecordPersonalInfo(ctx, input)
	if err != nil {
		return model.UserPersonalInfo{}, fmt.Errorf("error when record personal info for id='%d':%w", input.Id)
	}
	return user, nil
}
