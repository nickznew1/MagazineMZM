package item_repository_postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) DeleteItem(
	ctx context.Context,
	input model.Item) (model.Item, error) {

	var item model.Item
	var document model.ItemSpecFiles

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
    SELECT link 
    FROM item_spec_files 
    WHERE id = $1
    `

	row := r.pool.QueryRow(ctx, query, input.Id)

	err := row.Scan(
		&document.SpecFileLink,
	)
	if err != nil {
		return model.Item{}, fmt.Errorf("scan query error :%w", err)
	}

	err = os.Remove("./public/documents/" + document.SpecFileLink)
	if err != nil {
		return model.Item{}, fmt.Errorf("os.remove error :%w", err)
	}

	query = `
    SELECT 
    item_picture 
    FROM item 
    WHERE id =$1
    `

	row = r.pool.QueryRow(ctx, query, input.Id)

	err = row.Scan(
		&item.ItemPicture,
	)

	if err != nil {
		return model.Item{}, fmt.Errorf("scan row error: %w", err)
	}

	err = os.Remove("./public/images/" + item.ItemPicture)
	if err != nil {
		return model.Item{}, fmt.Errorf("os.Remove error: %w", err)
	}

	query = `
       DELETE FROM item 
       WHERE id =$1
       `

	cmpTag, err := r.pool.Exec(ctx, query, input.Id)
	if err != nil {
		return model.Item{}, fmt.Errorf("exec query error: %w", err)
	}
	if cmpTag.RowsAffected() == 0 {
		return model.Item{}, fmt.Errorf("no changes in database: %w", err)
	}

	return item, nil
}
