package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) GetItemId(
	ctx context.Context,
	input model.Item) (model.Item, error) {

	item, err := s.itemRepository.GetItemId(ctx, input)
	if err != nil {
		return model.Item{}, fmt.Errorf("failed to get item id :%w", err)
	}
	return item, err
}
