package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"time"
)

type ApplicationRepo struct {
	db *pgxpool.Pool
}

func NewApplicationRepo(db *pgxpool.Pool) model.ApplicationRepository {
	return &ApplicationRepo{
		db: db}
}

func (s *ApplicationRepo) CreateApplication(ctx context.Context, input model.Application) (string, error) {
	fmt.Println("create application")
	var applicationId string
	date := time.Now()
	jsonItems, err := json.Marshal(input.Items)
	if err != nil {
		return "", err
	}

	err = s.db.QueryRow(ctx, `INSERT INTO 
    user_applications(user_id, email, first_name, second_name, login, phone_number, company, address, city, order_date, items) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    RETURNING id`,
		input.UserId,
		input.Email,
		input.FirstName,
		input.SecondName,
		input.Login,
		input.PhoneNumber,
		input.Company,
		input.Address,
		input.City,
		date,
		jsonItems).Scan(&applicationId)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return applicationId, nil
}

func (s *ApplicationRepo) GetApplication(ctx context.Context, id string, userId string) (model.Application, error) {
	fmt.Println("get application")
	var application model.Application
	var temp string
	var selectedItems []byte
	var resultItems []model.Cart
	err := s.db.QueryRow(ctx, `SELECT * 
    from user_applications WHERE id=$1 AND user_id =$2`,
		id, userId).Scan(
		&temp,
		&application.UserId,
		&application.Email,
		&application.FirstName,
		&application.SecondName,
		&application.Login,
		&application.PhoneNumber,
		&application.Company,
		&application.Address,
		&application.City,
		&application.OrderDate,
		&selectedItems,
		&application.Status,
	)
	if err != nil {
		fmt.Println(err)
		return application, err
	}
	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		fmt.Println("cant unmarshal items")
		return application, err
	}
	_, err = s.db.Exec(ctx, "DELETE FROM customer_item WHERE customer_id = $1", userId)
	if err != nil {
		fmt.Println(err)
		return application, nil
	}
	application.Items = resultItems
	return application, nil
}

func (s *ApplicationRepo) GetAllApplicationsForUser(ctx context.Context, userId string) ([]model.Application, error) {
	fmt.Println("get applications")
	var applications []model.Application
	var allItems []byte
	var resultItems []model.Cart
	rows, err := s.db.Query(ctx, `SELECT *  
    from user_applications WHERE user_id =$1`, userId)
	if err != nil {
		fmt.Println(err)
		return applications, err
	}
	for rows.Next() {
		var application model.Application
		err = rows.Scan(
			&application.Id,
			&application.UserId,
			&application.Email,
			&application.FirstName,
			&application.SecondName,
			&application.Login,
			&application.PhoneNumber,
			&application.Company,
			&application.Address,
			&application.City,
			&application.OrderDate,
			&allItems,
			&application.Status,
		)
		if err != nil {
			fmt.Println(err)
			return applications, err
		}
		err = json.Unmarshal(allItems, &resultItems)
		if err != nil {
			fmt.Println("cant unmarshal items")
			return applications, err
		}
		application.Items = resultItems
		applications = append(applications, application)
	}

	return applications, nil
}

func (s *ApplicationRepo) GetAllApplicationsForAdmin(ctx context.Context) ([]model.Application, error) {
	var applications []model.Application
	var allItems []byte
	var resultItems []model.Cart
	rows, err := s.db.Query(ctx, `SELECT *  
    from user_applications ORDER BY id`)
	if err != nil {
		fmt.Println(err)
		return applications, err
	}
	for rows.Next() {
		var application model.Application
		err = rows.Scan(
			&application.Id,
			&application.UserId,
			&application.Email,
			&application.FirstName,
			&application.SecondName,
			&application.Login,
			&application.PhoneNumber,
			&application.Company,
			&application.Address,
			&application.City,
			&application.OrderDate,
			&allItems,
			&application.Status,
		)
		if err != nil {
			fmt.Println(err)
			return applications, err
		}
		err = json.Unmarshal(allItems, &resultItems)
		if err != nil {
			fmt.Println("cant unmarshal items")
			return applications, err
		}
		application.Items = resultItems
		applications = append(applications, application)
	}

	return applications, nil
}

func (s *ApplicationRepo) SetNewApplicationStatus(ctx context.Context, input model.Application) (model.Application, error) {
	var application model.Application
	fmt.Println(input, "ZDES")
	err := s.db.QueryRow(ctx, `UPDATE user_applications SET order_status =$1 WHERE id = $2 RETURNING order_status`, input.Status, input.Id).Scan(&application.Status)
	if err != nil {
		fmt.Println(err)
		return application, err
	}
	return application, nil
}

func (s *ApplicationRepo) GetApplicationForAdmin(ctx context.Context, id string) (model.Application, error) {
	fmt.Println("get application")
	var application model.Application
	var selectedItems []byte
	var resultItems []model.Cart
	err := s.db.QueryRow(ctx, `SELECT * 
    from user_applications WHERE id=$1`,
		id).Scan(
		&application.Id,
		&application.UserId,
		&application.Email,
		&application.FirstName,
		&application.SecondName,
		&application.Login,
		&application.PhoneNumber,
		&application.Company,
		&application.Address,
		&application.City,
		&application.OrderDate,
		&selectedItems,
		&application.Status,
	)
	if err != nil {
		fmt.Println(err)
		return application, err
	}
	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		fmt.Println("cant unmarshal items")
		return application, err
	}

	application.Items = resultItems
	return application, nil
}
