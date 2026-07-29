package usecase

import "myWebApp/backend/internal/domain/model"

type ApplicationUseCase struct {
	repo model.ApplicationRepository
}

func NewApplicationUseCase(r model.ApplicationRepository) *ApplicationUseCase {
	return &ApplicationUseCase{repo: r}
}

func (a *ApplicationUseCase) GetAllApplicationsForUser(userId string) ([]model.Application, error) {
	return a.repo.GetAllApplicationsForUser(userId)
}

func (a *ApplicationUseCase) GetApplication(id string, userId string) (model.Application, error) {
	return a.repo.GetApplication(id, userId)
}

func (a *ApplicationUseCase) CreateApplication(input model.Application) (string, error) {
	return a.repo.CreateApplication(input)
}

func (a *ApplicationUseCase) GetAllApplicationsForAdmin() ([]model.Application, error) {
	return a.repo.GetAllApplicationsForAdmin()
}

func (a *ApplicationUseCase) SetNewApplicationStatus(input model.Application) (model.Application, error) {
	return a.repo.SetNewApplicationStatus(input)
}

func (a *ApplicationUseCase) GetApplicationForAdmin(id string) (model.Application, error) {
	return a.repo.GetApplicationForAdmin(id)
}
