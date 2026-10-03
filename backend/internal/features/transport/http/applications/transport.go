package applications_transport_http

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type ApplicationHTTPHandler struct {
	applicationService ApplicationService
}

type ApplicationService interface {
	GetAllApplicationsForUser(
		ctx context.Context,
		userId string,
	) ([]model.Application, error)

	GetApplication(
		ctx context.Context,
		id string, userId string,
	) (model.Application, error)

	CreateApplication(
		ctx context.Context,
		input model.Application,
	) (string, error)

	GetAllApplicationsForAdmin(
		ctx context.Context,
	) ([]model.Application, error)

	SetApplicationStatus(
		ctx context.Context,
		input model.Application,
	) (model.Application, error)

	GetApplicationForAdmin(
		ctx context.Context,
		id string,
	) (model.Application, error)
}

func NewApplicationHTTPHandler(applicationService ApplicationService) *ApplicationHTTPHandler {
	return &ApplicationHTTPHandler{
		applicationService: applicationService,
	}
}
