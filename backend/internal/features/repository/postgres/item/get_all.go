package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) GetAll(
	ctx context.Context) ([]model.Item, error) {

	var items []model.Item

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
       SELECT *
       FROM item
       `

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return []model.Item{}, fmt.Errorf("query rows error :%w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&item.ItemType,
			&item.ItemSecondaryType,
			&item.ItemPicture,
			&item.ItemDescription,
			&item.ItemShortDescription,
			&item.Article,
			&item.Visible)
		if err != nil {
			return []model.Item{}, fmt.Errorf("rows scan error :%w", err)
		}
		items = append(items, item)
	}

	return items, nil
}
