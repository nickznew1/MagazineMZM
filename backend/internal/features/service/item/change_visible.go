package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) ChangeVisible(
	ctx context.Context,
	input bool,
	id string) (model.Item, error) {

	item, err := s.itemRepository.ChangeVisible(ctx, input, id)
	if err != nil {
		return model.Item{}, fmt.Errorf("failed to change visible for item with id ='%s' : %w", id, err)
	}
	return item, nil
}
