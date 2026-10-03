package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) SetApplicationStatus(
	ctx context.Context,
	input model.Application) (model.Application, error) {

	application, err := s.applicationRepository.SetApplicationStatus(ctx, input)
	if err != nil {
		return model.Application{}, fmt.Errorf("failed to change application id ='%d' status :%w", input.Id, err)
	}
	return application, nil
}
