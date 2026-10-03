package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) GetAllApplicationsForAdmin(
	ctx context.Context,
) ([]model.Application, error) {

	applications, err := s.GetAllApplicationsForAdmin(ctx)
	if err != nil {
		return []model.Application{}, fmt.Errorf("failed to get all applications for admin :%w", err)
	}
	return applications, nil
}
