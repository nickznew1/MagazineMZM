package usecase

import (
	"context"
	"fmt"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type UserUseCase struct {
	repo model.UserRepository
}

func NewUserUseCase(r model.UserRepository) *UserUseCase {

	fmt.Println("CREATING USECASE")
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.GetUserById(ctx, input)
}
func (u *UserUseCase) CreateUser(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.CreateUser(ctx, input)
}
func (u *UserUseCase) UserAuth(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.UserAuth(ctx, input)
}

func (u *UserUseCase) RecordPersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	return u.repo.RecordPersonalInfo(ctx, input)
}

func (u *UserUseCase) UpdatePersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	return u.repo.UpdatePersonalInfo(ctx, input)
}

func (u *UserUseCase) RecordDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	return u.repo.RecordDeliveryInfo(ctx, input)
}
func (u *UserUseCase) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	return u.repo.UpdateDeliveryInfo(ctx, input)
}
func (u *UserUseCase) UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error) {
	return u.repo.UserPasswordChange(ctx, input)
}
func (u *UserUseCase) UserChangeEmail(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.UserChangeEmail(ctx, input)
}

func (u *UserUseCase) GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error) {
	return u.repo.GetAllUsers(ctx)
}

func (u *UserUseCase) FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error) {
	return u.repo.FetchProfileInfo(ctx, id)
}

func (u *UserUseCase) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	return u.repo.FetchProfilePersonalInfo(ctx, id)
}

func (u *UserUseCase) FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error) {
	return u.repo.FetchProfileDeliveryInfo(ctx, id)
}
