package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) CreateItem(
	ctx context.Context,
	input model.Item,
	documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error) {

	item, documents, err := s.itemRepository.CreateItem(ctx, input, documents)
	if err != nil {
		return model.Item{}, model.ItemSpecFiles{}, fmt.Errorf("failed to create item: %w", err)
	}
	return item, documents, nil
}
