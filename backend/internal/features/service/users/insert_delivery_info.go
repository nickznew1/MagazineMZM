package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) RecordDeliveryInfo(
	ctx context.Context,
	input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {

	user, err := s.RecordDeliveryInfo(ctx, input)
	if err !=nil{
		return model.UserDeliveryInfo{}, fmt.Errorf("error when record delivery info for id='%d': %w", input.Id, err)
	}
	return user, nil
}
