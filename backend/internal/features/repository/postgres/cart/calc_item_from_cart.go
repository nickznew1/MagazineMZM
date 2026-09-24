package cart_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *CartRepository) CalcItemFromCart(
	ctx context.Context,
	input model.Cart) (model.Cart, error) {

	var calc model.Cart

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
  UPDATE customer_item 
  SET count =$1 
  WHERE item_id =$2 
  AND customer_id = $3 
  RETURNING item_id,count, customer_id
     `

	row := r.pool.QueryRow(ctx, query, input.Count, input.ItemId, input.Id)

	err := row.Scan(
		&calc.ItemSpecId,
		&calc.Count,
		&calc.Id,
	)

	if err != nil {
		return model.Cart{}, fmt.Errorf("scan query error: %w", err)
	}

	return calc, err
}
