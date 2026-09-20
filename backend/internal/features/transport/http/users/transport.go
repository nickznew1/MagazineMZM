package users_transport_http

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"github.com/nickznew1/MagazineMZM/backend/pkg/auth"
)

type UsersHTTPHandler struct {
	usersService UsersService
	auth         auth.TokenManager
	cartUseCase  *usecase.CartUseCase
}

type UsersService interface {
	GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error)
	CreateUser(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error)
	UserAuth(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error)
	FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error)
	FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error)
	FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error)
	RecordPersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error)
	RecordDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error)
	UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error)
	UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error)
	UserChangeEmail(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error)
	GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error)
}

func NewUsersHTTPHandler(usersService UsersService, auth auth.TokenManager) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
		auth:         auth,
	}
}
