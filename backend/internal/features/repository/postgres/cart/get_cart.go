package cart_repository_postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *CartRepository) GetCart(
	ctx context.Context,
	id string) ([]model.Cart, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	var shoppingCart []model.Cart

	query := `
    SELECT
    item_id,
    count,
    customer_id, 
    props,
    item.name, 
    item.price, 
    item.item_picture, 
    item.item_description,
    item.article,
    item.item_type,
    item.secondary_type
    FROM customer_item 
        JOIN item ON customer_item.item_id = item.id 
        WHERE customer_id =$1`

	row, err := r.pool.Query(ctx, query, id)
	if err != nil {
		return []model.Cart{}, fmt.Errorf("query row error: %w", err)
	}

	for row.Next() {
		var cart model.Cart
		var raw []byte
		var props map[string]any
		err = row.Scan(
			&cart.ItemId,
			&cart.Count,
			&cart.Id,
			&raw,
			&cart.Name,
			&cart.Price,
			&cart.ItemPicture,
			&cart.ItemDescription,
			&cart.Article,
			&cart.ItemType,
			&cart.SecondaryItemType)
		if err != nil {
			return []model.Cart{}, fmt.Errorf("scan row error: %w", err)
		}
		err = json.Unmarshal(raw, &props)
		cart.Props = props
		shoppingCart = append(shoppingCart, cart)
	}

	return shoppingCart, nil
}
