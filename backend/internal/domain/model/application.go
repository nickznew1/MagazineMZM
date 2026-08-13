package model

import (
	"context"
	"time"
)

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
	GetAllApplicationsForUser(ctx context.Context, userId string) ([]Application, error)
	GetApplication(ctx context.Context, id string, userId string) (Application, error)
	CreateApplication(ctx context.Context, input Application) (string, error)
	GetAllApplicationsForAdmin(ctx context.Context) ([]Application, error)
	SetNewApplicationStatus(ctx context.Context, input Application) (Application, error)
	GetApplicationForAdmin(ctx context.Context, id string) (Application, error)
}
