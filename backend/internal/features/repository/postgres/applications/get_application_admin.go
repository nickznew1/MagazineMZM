package applications_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ApplicationRepository) GetAllApplicationsForAdmin(
	ctx context.Context,
) ([]model.Application, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var applications []model.Application
	var allItems []byte
	var resultItems []model.Cart

	query := `
    SELECT *  
    from user_applications 
    ORDER BY id
    `

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return []model.Application{}, fmt.Errorf("query error :%w", err)
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
			return []model.Application{}, fmt.Errorf("scan rows error: %w", err)
		}
		err = json.Unmarshal(allItems, &resultItems)
		if err != nil {
			return []model.Application{}, fmt.Errorf("json unmarshal error :%w", err)
		}
		application.Items = resultItems
		applications = append(applications, application)
	}

	return applications, nil
}
