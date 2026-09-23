package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error) {
	return s.FetchProfileDeliveryInfo(ctx, id)
}
