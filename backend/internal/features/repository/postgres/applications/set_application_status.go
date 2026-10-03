package applications_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ApplicationRepository) SetApplicationStatus(
	ctx context.Context,
	input model.Application) (model.Application, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var application model.Application

	query := `
    UPDATE user_applications 
    SET order_status =$1 
    WHERE id = $2 
    RETURNING order_status`

	row := r.pool.QueryRow(
		ctx,
		query,
		input.Status,
		input,
	)

	err := row.Scan(
		&application.Status,
	)

	if err != nil {
		return model.Application{}, fmt.Errorf("scan query error: %w", err)
	}

	return application, nil
}
