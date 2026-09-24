package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	user, err := s.UpdateDeliveryInfo(ctx, input)
	if err != nil {
		return model.UserDeliveryInfo{}, fmt.Errorf("error when update delivery info for id='%d': %w", input.Id, err)
	}
	return user, nil
}
