package core_domain_user

import "time"

type UserOrdinaryInfo struct {
	Id               int
	Login            string
	Password         string
	Email            string
	UserRole         string
	RegistrationDate time.Time
}

type UserOrdinaryInfoOut struct {
	Id       int
	Login    string
	Email    string
	UserRole string
}

type UserPersonalInfo struct {
	Id         int
	Company    string
	FirstName  string
	SecondName string
}

type UserPersonalInfoOut struct {
	Id         int
	Company    string
	FirstName  string
	SecondName string
}

type UserDeliveryInfo struct {
	Id          int
	PhoneNumber string
	City        string
	Address     string
}

type UserDeliveryInfoOut struct {
	Id          int
	PhoneNumber string
	City        string
	Address     string
}

type PasswordChange struct {
	OldPassword string
	NewPassword string
	Id          int
}

type UserMerge struct {
	Error error
	Data  interface{}
	Kind  string
}
