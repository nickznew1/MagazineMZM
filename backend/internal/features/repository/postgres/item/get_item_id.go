package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) GetItemId(
	ctx context.Context,
	input model.Item) (model.Item, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var item model.Item

	query := `
      SELECT id 
      FROM item 
      WHERE name =$1
    `
	rows := r.pool.QueryRow(ctx, query, input.Name)

	err := rows.Scan(
		&item.Id,
	)

	if err != nil {
		return model.Item{}, fmt.Errorf("scan query error: %w", err)
	}

	return item, nil
}
