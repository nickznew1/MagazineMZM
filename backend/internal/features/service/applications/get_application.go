package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) GetApplication(
	ctx context.Context,
	id string,
	userId string) (model.Application, error) {

	application, err := s.GetApplication(ctx, id, userId)
	if err != nil {
		return model.Application{}, fmt.Errorf("failed to get application id ='%s' for user id ='%s' : %w", id, userId, err)
	}
	return application, nil
}
