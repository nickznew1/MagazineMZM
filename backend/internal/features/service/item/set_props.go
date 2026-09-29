package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) SetProps(
	ctx context.Context,
	input []model.ItemProp,
	id string) ([]model.ItemProp, error) {

	props, err := s.itemRepository.SetProps(ctx, input, id)
	if err != nil {
		return []model.ItemProp{}, fmt.Errorf("failed to set props for item id ='%s' : %w", id, err)
	}
	return props, err
}
