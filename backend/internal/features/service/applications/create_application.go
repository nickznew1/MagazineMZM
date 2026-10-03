package applications_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ApplicationService) CreateApplication(
	ctx context.Context,
	input model.Application) (string /*NEED TO FIX*/, error) {

	application, err := s.applicationRepository.CreateApplication(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to create application: %w", err)
	}
	return application, nil
}
