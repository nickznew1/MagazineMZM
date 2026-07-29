package model

import "time"

type Application struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Email       string    `json:"email"`
	FirstName   string    `json:"first_name"`
	SecondName  string    `json:"second_name"`
	Login       string    `json:"login"`
	PhoneNumber string    `json:"phone_number"`
	Company     string    `json:"company"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	OrderDate   time.Time `json:"order_date"`
	Items       []Cart    `json:"items"`
	Status      string    `json:"application_status"`
}

type ApplicationRepository interface {
	GetAllApplicationsForUser(userId string) ([]Application, error)
	GetApplication(id string, userId string) (Application, error)
	CreateApplication(input Application) (string, error)
	GetAllApplicationsForAdmin() ([]Application, error)
	SetNewApplicationStatus(input Application) (Application, error)
	GetApplicationForAdmin(id string) (Application, error)
}
