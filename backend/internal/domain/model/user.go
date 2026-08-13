package model

import (
	"context"
	"time"
)

type UserOrdinaryInfo struct {
	Id               int       `json:"id"`
	Login            string    `json:"login"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	UserRole         string    `json:"user_role"`
	RegistrationDate time.Time `json:"registration_date"`
}

type UserOrdinaryInfoOut struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	UserRole string `json:"user_role"`
}

type UserPersonalInfo struct {
	Id         int    `json:"id"`
	Company    string `json:"company"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type UserPersonalInfoOut struct {
	Id         int    `json:"id"`
	Company    string `json:"company"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

type UserDeliveryInfo struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Address     string `json:"address"`
}

type UserDeliveryInfoOut struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	Address     string `json:"address"`
}

type UserSummary struct {
	UserOrdinary *UserOrdinaryInfoOut `json:"user_ordinary"`
	UserPersonal *UserPersonalInfoOut `json:"user_personal"`
	UserDelivery *UserDeliveryInfoOut `json:"user_delivery"`
}

type PasswordChange struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Id          int    `json:"id"`
}

type UserMerge struct {
	Error error
	Data  interface{}
	Kind  string
}

type UserRepository interface {
	GetUserById(ctx context.Context, input UserOrdinaryInfo) (UserOrdinaryInfo, error)
	CreateUser(ctx context.Context, input UserOrdinaryInfo) (UserOrdinaryInfo, error)
	UserAuth(ctx context.Context, input UserOrdinaryInfo) (UserOrdinaryInfo, error)
	FetchProfileInfo(ctx context.Context, id string) (UserOrdinaryInfoOut, error)
	FetchProfilePersonalInfo(ctx context.Context, id string) (UserPersonalInfoOut, error)
	FetchProfileDeliveryInfo(ctx context.Context, id string) (UserDeliveryInfoOut, error)
	RecordPersonalInfo(ctx context.Context, input UserPersonalInfo) (UserPersonalInfo, error)
	UpdatePersonalInfo(ctx context.Context, input UserPersonalInfo) (UserPersonalInfo, error)
	RecordDeliveryInfo(ctx context.Context, input UserDeliveryInfo) (UserDeliveryInfo, error)
	UpdateDeliveryInfo(ctx context.Context, input UserDeliveryInfo) (UserDeliveryInfo, error)
	UserPasswordChange(ctx context.Context, input PasswordChange) (PasswordChange, error)
	UserChangeEmail(ctx context.Context, input UserOrdinaryInfo) (UserOrdinaryInfo, error)
	GetAllUsers(ctx context.Context) ([]UserOrdinaryInfo, error)
}
