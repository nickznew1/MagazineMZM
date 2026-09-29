package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) GetAllPropsName(
	ctx context.Context) (model.ItemProp, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	var props model.ItemProp

	temp := make([]string, 30)
	var index int

	query := `
    SELECT * 
    FROM item_properties
    `

	rows, err := r.pool.Query(ctx, query)

	if err != nil {
		return model.ItemProp{}, fmt.Errorf("query rows error: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var prop string
		err = rows.Scan(&index, &prop)
		if err != nil {
			return model.ItemProp{}, fmt.Errorf("scan query error; %w", err)
		}
		temp[index] = prop
	}

	props.PropNameA = temp[:index]

	return props, nil
}
