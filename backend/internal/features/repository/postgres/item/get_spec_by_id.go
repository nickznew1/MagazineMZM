package item_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *ItemRepository) GetSpecById(
	ctx context.Context,
	id int) ([]model.ItemSpecFiles, error) {

	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var specStore []model.ItemSpecFiles

	query := `
    SELECT name, link, picture 
    FROM item_spec_files 
    WHERE id =$1
    `
	rows, err := r.pool.Query(ctx, query, id)

	if err != nil {
		return []model.ItemSpecFiles{}, fmt.Errorf("query rows error: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var specRow model.ItemSpecFiles
		err = rows.Scan(&specRow.SpecFileName,
			&specRow.SpecFileLink,
			&specRow.SpecFilePic)
		if err != nil {
			return []model.ItemSpecFiles{}, fmt.Errorf("scan rows error: %w", err)
		}
		specStore = append(specStore, specRow)
	}

	return specStore, nil
}
