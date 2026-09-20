package users_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	GetUserById(ctx context.Context,
		input model.UserOrdinaryInfo,
	) (model.UserOrdinaryInfo, error)

	CreateUser(ctx context.Context,
		input model.UserOrdinaryInfo,
	) (model.UserOrdinaryInfo, error)

	UserAuth(ctx context.Context,
		input model.UserOrdinaryInfo,
	) (model.UserOrdinaryInfo, error)

	FetchProfileInfo(ctx context.Context,
		id string,
	) (model.UserOrdinaryInfoOut, error)

	FetchProfilePersonalInfo(ctx context.Context,
		id string,
	) (model.UserPersonalInfoOut, error)

	FetchProfileDeliveryInfo(ctx context.Context,
		id string,
	) (model.UserDeliveryInfoOut, error)

	RecordPersonalInfo(ctx context.Context,
		input model.UserPersonalInfo,
	) (model.UserPersonalInfo, error)

	UpdatePersonalInfo(ctx context.Context,
		input model.UserPersonalInfo,
	) (model.UserPersonalInfo, error)

	RecordDeliveryInfo(ctx context.Context,
		input model.UserDeliveryInfo,
	) (model.UserDeliveryInfo, error)

	UpdateDeliveryInfo(ctx context.Context,
		input model.UserDeliveryInfo,
	) (model.UserDeliveryInfo, error)

	UserPasswordChange(ctx context.Context,
		input model.PasswordChange,
	) (model.PasswordChange, error)

	UserChangeEmail(ctx context.Context,
		input model.UserOrdinaryInfo,
	) (model.UserOrdinaryInfo, error)

	GetAllUsers(ctx context.Context,
	) ([]model.UserOrdinaryInfo, error)
}

func (s *UsersService) NewUsersService(usersRepository UsersRepository,
) *UsersService {

	return &UsersService{
		usersRepository: usersRepository,
	}
}
