package applications_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ApplicationRepository) GetApplicationForAdmin(
	ctx context.Context,
	id string) (model.Application, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var application model.Application
	var selectedItems []byte
	var resultItems []model.Cart
	query := `
    SELECT * 
    from user_applications 
    WHERE id=$1
    `

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
	)

	err := row.Scan(
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
		return model.Application{},
			fmt.Errorf("scan query error: %w", err)
	}

	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		return model.Application{}, fmt.Errorf("json unmarshal error: %w", err)
	}

	application.Items = resultItems

	return application, nil
}
