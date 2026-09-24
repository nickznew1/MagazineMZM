package cart_repository_postgres

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *CartRepository) DeleteFromCart(
	ctx context.Context,
	input model.Cart) ([]model.Cart, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
      DELETE FROM customer_item 
             WHERE item_id = $1 
               AND customer_id = $2
    `

	cmpTag, err := r.pool.Exec(ctx, query, input.ItemId, input.Id)
	if err != nil {
		return []model.Cart{}, fmt.Errorf("exec query error: %w", err)
	}
	if cmpTag.RowsAffected() == 0 {
		return []model.Cart{}, fmt.Errorf("cart with id='%d' not found:%w", input.Id, err)
	}

	newCart, err := r.GetCart(ctx, strconv.Itoa(input.Id))
	if err != nil {
		return []model.Cart{}, fmt.Errorf("error when getting new cart from repo: %w", err)
	}

	return newCart, nil
}
