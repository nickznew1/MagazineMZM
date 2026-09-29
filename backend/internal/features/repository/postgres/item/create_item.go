package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) CreateItem(
	ctx context.Context,
	input model.Item,
	documents model.ItemSpecFiles) (
	model.Item,
	model.ItemSpecFiles, error) {

	var item model.Item
	var document model.ItemSpecFiles

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `INSERT INTO item 
    (name,price,item_type, secondary_type, item_picture, item_description, item_short_description, article) 
    VALUES ($1,$2,$3,$4,$5,$6,$7, $8) 
    RETURNING id,name,price,item_type`

	rows := r.pool.QueryRow(ctx, query, input.Name, input.Price, input.ItemType, input.ItemSecondaryType, input.ItemPicture, input.ItemDescription, input.ItemShortDescription, input.Article)

	err := rows.Scan(
		&item.Id,
		&item.Name,
		&item.Price,
		&item.ItemType)

	if err != nil {
		return model.Item{}, model.ItemSpecFiles{}, fmt.Errorf("scan query error: %w", err)
	}

	query = `
    INSERT INTO item_spec_files 
    (id,name,link,picture) 
    VALUES ($1, $2, $3,$4)
    `
	cmpTag, err := r.pool.Exec(ctx, query, item.Id, documents.SpecFileName, documents.SpecFileLink, documents.SpecFilePic)

	if err != nil {
		return model.Item{}, model.ItemSpecFiles{}, fmt.Errorf("exec query error: %w", err)
	}
	if cmpTag.RowsAffected() == 0 {
		return model.Item{}, model.ItemSpecFiles{}, fmt.Errorf("no changes in database: %w", err)
	}
	return item, document, nil
}
