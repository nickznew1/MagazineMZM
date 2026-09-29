package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) DeleteItem(
	ctx context.Context,
	input model.Item) (model.Item, error) {

	item, err := s.itemRepository.DeleteItem(ctx, input)
	if err != nil {
		return model.Item{}, fmt.Errorf("failed to delete item with id ='%d' : %w", input.Id, err)
	}
	return item, nil
}
