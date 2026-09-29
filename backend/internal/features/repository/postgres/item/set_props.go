package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) SetProps(
	ctx context.Context,
	input []model.ItemProp,
	id string) ([]model.ItemProp, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var props []model.ItemProp

	idSlice := make([]int, 0, len(input))
	valueSlice := make([]string, 0, len(input))

	for _, item := range input {
		idSlice = append(idSlice, item.IdProp)
		valueSlice = append(valueSlice, item.PropValue)
	}

	query := `
    INSERT INTO item_properties_values 
    (item_id,property_id, value) 
    SELECT $1, unnest($2::int[]), unnest($3::text[]) `

	cmpTag, err := r.pool.Exec(ctx, query, id, idSlice, valueSlice)
	if err != nil {
		return []model.ItemProp{}, fmt.Errorf("exec query error: %w", err)
	}
	if cmpTag.RowsAffected() == 0 {
		return []model.ItemProp{}, fmt.Errorf("no changes in database: %w", err)
	}

	return props, nil
}
