package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) GetSpecById(
	ctx context.Context,
	id int) ([]model.ItemSpecFiles, error) {

	spec, err := s.itemRepository.GetSpecById(ctx, id)
	if err != nil {
		return []model.ItemSpecFiles{}, fmt.Errorf("failed to get spec for item with id='%d' : %w", id, err)
	}
	return spec, nil

}
