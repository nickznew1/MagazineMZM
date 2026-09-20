package users_repository_postgres

import "time"

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
