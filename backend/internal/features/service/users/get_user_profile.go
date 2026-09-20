package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error) {
	return s.FetchProfileInfo(ctx, id)
}
