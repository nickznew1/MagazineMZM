package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	user, err := s.usersRepository.FetchProfilePersonalInfo(ctx, id)
	if err != nil {
		return model.UserPersonalInfoOut{}, fmt.Errorf("error when fetch profile personal info with id='%s': %w", id, err)
	}
	return user, nil
}
