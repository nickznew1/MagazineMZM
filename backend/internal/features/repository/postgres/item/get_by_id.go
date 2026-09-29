package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) GetById(
	ctx context.Context,
	id int) (model.ItemProp, error) {

	var props model.ItemProp

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `SELECT 
    item.id,
    item.name, 
    item.price, 
    item.item_description, 
    item.item_short_description,
    item.item_picture,
    item.article,
    item.item_type,
    item.secondary_type,
    item.visible
    FROM item WHERE item.id =$1`

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(
		&props.Id,
		&props.Name,
		&props.Price,
		&props.ItemDescription,
		&props.ItemShortDescription,
		&props.ItemPicture,
		&props.Article,
		&props.ItemType,
		&props.ItemSecondaryType,
		&props.Visible)
	if err != nil {
		return model.ItemProp{}, fmt.Errorf("scan query error : %w", err)
	}

	query = `SELECT 
    item_properties.name,
    item_properties_values.value 
    FROM item 
    JOIN item_properties_values ON item.id = item_properties_values.item_id 
        JOIN item_properties ON item_properties_values.property_id = item_properties.id 
    WHERE item.id =$1`

	rows, err := r.pool.Query(ctx, query, id)
	if err != nil {
		return model.ItemProp{}, fmt.Errorf("query row error: %w", err)
	}
	defer rows.Close()
	for rows.Next() {

		var name string
		var prop string

		err = rows.Scan(&name, &prop)
		if err != nil {
			return model.ItemProp{}, fmt.Errorf("scan query error: %w", err)
		}
		props.PropNameA = append(props.PropNameA, name)
		props.PropValueA = append(props.PropValueA, prop)
	}

	return props, nil
}
