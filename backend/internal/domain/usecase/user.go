package usecase

import (
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

func (u *UserUseCase) GetUserById(input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.GetUserById(input)
}
func (u *UserUseCase) CreateUser(input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.CreateUser(input)
}
func (u *UserUseCase) UserAuth(input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.UserAuth(input)
}

func (u *UserUseCase) RecordPersonalInfo(input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	return u.repo.RecordPersonalInfo(input)
}

func (u *UserUseCase) UpdatePersonalInfo(input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	return u.repo.UpdatePersonalInfo(input)
}

func (u *UserUseCase) RecordDeliveryInfo(input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	return u.repo.RecordDeliveryInfo(input)
}
func (u *UserUseCase) UpdateDeliveryInfo(input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	return u.repo.UpdateDeliveryInfo(input)
}
func (u *UserUseCase) UserPasswordChange(input model.PasswordChange) (model.PasswordChange, error) {
	return u.repo.UserPasswordChange(input)
}
func (u *UserUseCase) UserChangeEmail(input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	return u.repo.UserChangeEmail(input)
}

func (u *UserUseCase) GetAllUsers() ([]model.UserOrdinaryInfo, error) {
	return u.repo.GetAllUsers()
}

func (u *UserUseCase) FetchProfileInfo(id string) (model.UserOrdinaryInfoOut, error) {
	return u.repo.FetchProfileInfo(id)
}

func (u *UserUseCase) FetchProfilePersonalInfo(id string) (model.UserPersonalInfoOut, error) {
	return u.repo.FetchProfilePersonalInfo(id)
}

func (u *UserUseCase) FetchProfileDeliveryInfo(id string) (model.UserDeliveryInfoOut, error) {
	return u.repo.FetchProfileDeliveryInfo(id)
}
