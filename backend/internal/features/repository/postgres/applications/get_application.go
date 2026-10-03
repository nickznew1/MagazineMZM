package applications_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ApplicationRepository) GetApplication(
	ctx context.Context,
	id string,
	userId string) (model.Application, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var application model.Application
	var temp string
	var selectedItems []byte
	var resultItems []model.Cart

	query := `
    SELECT * 
    from user_applications 
    WHERE id=$1 AND user_id =$2
    `

	row := r.pool.QueryRow(
		ctx,
		query,
		id,
		userId,
	)

	err := row.Scan(
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
		return model.Application{},
			fmt.Errorf("scan query error :%w", err)
	}

	err = json.Unmarshal(selectedItems, &resultItems)
	if err != nil {
		return model.Application{},
			fmt.Errorf("json unmarshal error: %w", err)
	}

	query = `
    DELETE FROM customer_item 
           WHERE customer_id = $1`

	cmpTag, err := r.pool.Exec(ctx, query)
	if err != nil {
		return model.Application{},
			fmt.Errorf("exec query error: %w", err)
	}
	if cmpTag.RowsAffected() == 0 {
		return model.Application{},
			fmt.Errorf("no rows affected :%w", err)
	}
	application.Items = resultItems

	return application, nil
}
