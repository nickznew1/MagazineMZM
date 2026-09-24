package cart_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *CartRepository) AddToCart(
	ctx context.Context,
	input model.Cart) (model.Cart, error) {
	var userItem model.Cart
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query :=
		`
        INSERT INTO customer_item(customer_id, count, item_id, props) 
        VALUES($1, $2, $3, $4) 
        RETURNING customer_id, count, item_id
    `

	propsJson, _ := json.Marshal(input.Props)

	row := r.pool.QueryRow(ctx, query, input.Id, input.Count, input.ItemId, propsJson)

	err := row.Scan(
		&userItem.Id,
		&userItem.Count,
		&userItem.ItemId,
	)

	if err != nil {
		return model.Cart{}, fmt.Errorf("scan query error: %w", err)
	}

	return userItem, nil
}
