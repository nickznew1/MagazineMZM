package users_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *UsersService) FetchProfileInfo(
	ctx context.Context,
	id string) (model.UserOrdinaryInfoOut, error) {

	user, err := s.FetchProfileInfo(ctx, id)
	if err != nil {
		return model.UserOrdinaryInfoOut{}, fmt.Errorf("error when fetch profile info with id='%s': %w", id, err)
	}
	return user, nil
}
