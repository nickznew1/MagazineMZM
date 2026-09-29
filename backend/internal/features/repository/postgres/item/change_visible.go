package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) ChangeVisible(
	ctx context.Context,
	status bool,
	id string) (model.Item, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var visible model.Item

	query := `
    UPDATE item 
    SET visible = $1 
    WHERE id = $2 
    RETURNING visible`

	row := r.pool.QueryRow(ctx, query, status, id)

	err := row.Scan(
		&visible.Visible,
	)

	if err != nil {
		return model.Item{}, fmt.Errorf("scan query error: %w", err)
	}

	return visible, nil
}
