package model

import "time"

type UserOrdinaryInfo struct {
	Id               int       `json:"id"`
	Login            string    `json:"login"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	UserRole         string    `json:"user_role"`
	RegistrationDate time.Time `json:"registration_date"`
}

type UserPersonalInfo struct {
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

type UserSummary struct {
	UserOrdinary UserOrdinaryInfo `json:"user_ordinary"`
	UserPersonal UserPersonalInfo `json:"user_personal"`
	UserDelivery UserDeliveryInfo `json:"user_delivery"`
}

type PasswordChange struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Id          int    `json:"id"`
}

type UserMerge struct {
	error error
	data  interface{}
	kind  string
}

type UserRepository interface {
	GetUserById(input User) (User, error)
	CreateUser(input User) (User, error)
	UserAuth(input User) (User, error)
	FetchProfile(id string) (UserSummary, error)
	FetchProfileInfo(id string) (UserOrdinaryInfo, error)
	FetchProfilePersonalInfo(id string) (UserPersonalInfo, error)
	FetchProfileDeliveryInfo(id string) (UserDeliveryInfo, error)
	RecordPersonalInfo(input UserPersonalInfo) (UserPersonalInfo, error)
	UpdatePersonalInfo(input UserPersonalInfo) (UserPersonalInfo, error)
	RecordDeliveryInfo(input UserDeliveryInfo) (UserDeliveryInfo, error)
	UpdateDeliveryInfo(input UserDeliveryInfo) (UserDeliveryInfo, error)
	UserPasswordChange(input PasswordChange) (PasswordChange, error)
	UserChangeEmail(input User) (User, error)
	GetAllUsers() ([]User, error)
}
