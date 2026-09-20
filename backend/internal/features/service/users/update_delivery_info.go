package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	return s.UpdateDeliveryInfo(ctx, input)
}
