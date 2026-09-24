package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfileDeliveryInfo(
	ctx context.Context,
	id string) (model.UserDeliveryInfoOut, error) {

	user, err := s.FetchProfileDeliveryInfo(ctx, id)
	if err != nil {
		return model.UserDeliveryInfoOut{}, fmt.Errorf("error when fetching profile delivery info(service) with id ='%s': %w", id, err)
	}
	return user, nil
}
