package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) GetApplicationForAdmin(
	ctx context.Context,
	id string) (model.Application, error) {

	application, err := s.applicationRepository.GetApplicationForAdmin(ctx, id)
	if err != nil {
		return model.Application{}, fmt.Errorf("failed to get application id ='%s' for admin :%w", id, err)
	}
	return application, nil
}
