package repository

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"time"
)

type ApplicationRepo struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewApplicationRepo(db *pgxpool.Pool, logger *slog.Logger) model.ApplicationRepository {
	return &ApplicationRepo{
		logger: logger,
		db:     db}
}

func (s *ApplicationRepo) CreateApplication(ctx context.Context, input model.Application) (string, error) {
	s.logger.Debug("Repository: CreateApplication started", "input: ", input)
	var applicationId string
	date := time.Now()
	jsonItems, err := json.Marshal(input.Items)
	if err != nil {
		s.logger.Error("Repository: CreateApplication error when marshal items to json", slog.Any("json_error: ", err))
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
		s.logger.Error("Repository: CreateApplication error when trying to append items to a result slice", slog.Any("db_error: ", err))
		return "", err
	}
	s.logger.Debug("Repository: CreateApplication success", "application_id: ", applicationId)
	return applicationId, nil
}

func (s *ApplicationRepo) GetApplication(ctx context.Context, id string, userId string) (model.Application, error) {
	s.logger.Debug("Repository: GetApplication for user started", "application_id: ", id, "user_id: ", userId)
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
		s.logger.Error("Repository: GetApplication for user error when trying to append items to a result slice", slog.Any("db_error: ", err))
		return application, err
	}
	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		s.logger.Error("Repository: GetApplication for user error when unmarshal selected items to a result slice", slog.Any("json_error: ", err))
		return application, err
	}
	_, err = s.db.Exec(ctx, "DELETE FROM customer_item WHERE customer_id = $1", userId)
	if err != nil {
		s.logger.Error("Repository: GetApplication for user error when delete application from user_cart", slog.Any("db_error: ", err))
		return application, nil
	}
	application.Items = resultItems
	s.logger.Debug("Repository: GetApplication for user success", "application_id: ", id, "user_id: ", userId)
	return application, nil
}

func (s *ApplicationRepo) GetAllApplicationsForUser(ctx context.Context, userId string) ([]model.Application, error) {
	s.logger.Debug("Repository: GetAllApplicationsForUser started", "user_id: ", userId)
	var applications []model.Application
	var allItems []byte
	var resultItems []model.Cart
	rows, err := s.db.Query(ctx, `SELECT *  
    from user_applications WHERE user_id =$1`, userId)
	if err != nil {
		s.logger.Error("Repository: GetAllApplicationsForUser error when select all from user applications", slog.Any("db_error: ", err))
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
			s.logger.Error("Repository: GetAllApplicationsForUser error when append applications to a result slice", slog.Any("db_error: ", err))
			return applications, err
		}
		err = json.Unmarshal(allItems, &resultItems)
		if err != nil {
			s.logger.Error("Repository: GetAllApplicationsForUser error when unmarshal json", slog.Any("json_error: ", err))
			return applications, err
		}
		application.Items = resultItems
		applications = append(applications, application)
	}
	s.logger.Debug("Repository: GetAllApplicationsForUser success", "user_id: ", userId)
	return applications, nil
}

func (s *ApplicationRepo) GetAllApplicationsForAdmin(ctx context.Context) ([]model.Application, error) {
	var applications []model.Application
	var allItems []byte
	var resultItems []model.Cart
	s.logger.Debug("Repository: GetAllApplicationsForAdmin started")
	rows, err := s.db.Query(ctx, `SELECT *  
    from user_applications ORDER BY id`)
	if err != nil {
		s.logger.Error("Repository: GetAllApplicationsForAdmin error when select all applications of user for admin", slog.Any("db_error: ", err))
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
			s.logger.Error("Repository: GetAllApplicationsForAdmin error when append applications from db to a result slice", slog.Any("db_error: ", err))
			return applications, err
		}
		err = json.Unmarshal(allItems, &resultItems)
		if err != nil {
			s.logger.Error("Repository: GetAllApplicationsForAdmin error when unmarshal all items to a result slice", slog.Any("json_error: ", err))
			return applications, err
		}
		application.Items = resultItems
		applications = append(applications, application)
	}
	s.logger.Debug("Repository: GetAllApplicationsForAdmin success")
	return applications, nil
}

func (s *ApplicationRepo) SetNewApplicationStatus(ctx context.Context, input model.Application) (model.Application, error) {
	var application model.Application
	s.logger.Debug("Repository: SetNewApplicationStatus started", "input: ", input)
	err := s.db.QueryRow(ctx, `UPDATE user_applications SET order_status =$1 WHERE id = $2 RETURNING order_status`, input.Status, input.Id).Scan(&application.Status)
	if err != nil {
		s.logger.Error("Repository: SetNewApplicationStatus error when trying to update order status in db", slog.Any("db_error: ", err))
		return application, err
	}
	s.logger.Debug("Repository: SetNewApplicationStatus success", "input: ", input)
	return application, nil
}

func (s *ApplicationRepo) GetApplicationForAdmin(ctx context.Context, id string) (model.Application, error) {
	s.logger.Debug("Repository: GetApplicationForAdmin started", "id: ", id)
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
		s.logger.Error("Repository: GetApplicationForAdmin error when trying to append items to a result slice", slog.Any("db_error: ", err))
		return application, err
	}
	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		s.logger.Error("Repository: GetApplicationForAdmin error when trying to unmarshal selected items to a result slice", slog.Any("json_error: ", err))
		return application, err
	}

	application.Items = resultItems
	s.logger.Debug("Repository: GetApplicationForAdmin success", "id: ", id)
	return application, nil
}
