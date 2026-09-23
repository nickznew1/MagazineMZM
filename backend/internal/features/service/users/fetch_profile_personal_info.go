package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	return s.usersRepository.FetchProfilePersonalInfo(ctx, id)
}
