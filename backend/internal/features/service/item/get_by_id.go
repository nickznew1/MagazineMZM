package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) GetById(
	ctx context.Context,
	id int) (model.ItemProp, error) {

	item, err := s.itemRepository.GetById(ctx, id)
	if err != nil {
		return model.ItemProp{}, fmt.Errorf("failed to get item with id='%d' : %w", id, err)
	}
	return item, err
}
