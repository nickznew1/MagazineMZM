package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) GetAll(
	ctx context.Context) ([]model.Item, error) {

	items, err := s.itemRepository.GetAll(ctx)
	if err != nil {
		return []model.Item{}, fmt.Errorf("failed to get all items: %w", err)
	}
	return items, nil
}
