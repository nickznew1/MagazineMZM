package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return s.GetUserById(ctx, input)
}
