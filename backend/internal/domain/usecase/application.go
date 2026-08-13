package usecase

import (
	"context"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type ApplicationUseCase struct {
	repo model.ApplicationRepository
}

func NewApplicationUseCase(r model.ApplicationRepository) *ApplicationUseCase {
	return &ApplicationUseCase{repo: r}
}

func (a *ApplicationUseCase) GetAllApplicationsForUser(ctx context.Context, userId string) ([]model.Application, error) {
	return a.repo.GetAllApplicationsForUser(ctx, userId)
}

func (a *ApplicationUseCase) GetApplication(ctx context.Context, id string, userId string) (model.Application, error) {
	return a.repo.GetApplication(ctx, id, userId)
}

func (a *ApplicationUseCase) CreateApplication(ctx context.Context, input model.Application) (string, error) {
	return a.repo.CreateApplication(ctx, input)
}

func (a *ApplicationUseCase) GetAllApplicationsForAdmin(ctx context.Context) ([]model.Application, error) {
	return a.repo.GetAllApplicationsForAdmin(ctx)
}

func (a *ApplicationUseCase) SetNewApplicationStatus(ctx context.Context, input model.Application) (model.Application, error) {
	return a.repo.SetNewApplicationStatus(ctx, input)
}

func (a *ApplicationUseCase) GetApplicationForAdmin(ctx context.Context, id string) (model.Application, error) {
	return a.repo.GetApplicationForAdmin(ctx, id)
}
