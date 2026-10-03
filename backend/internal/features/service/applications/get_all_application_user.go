package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) GetAllApplicationsForUser(
	ctx context.Context,
	userId string) ([]model.Application, error) {

	applications, err := s.applicationRepository.GetAllApplicationsForUser(ctx, userId)
	if err != nil {
		return []model.Application{},
			fmt.Errorf("failed to get all applications for user with id ='%s' :%w", userId, err)
	}
	return applications, err
}
