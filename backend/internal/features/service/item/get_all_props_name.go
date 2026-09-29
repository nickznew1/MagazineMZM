package item_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *ItemService) GetAllPropsName(
	ctx context.Context) (model.ItemProp, error) {

	props, err := s.itemRepository.GetAllPropsName(ctx)
	if err != nil {
		return model.ItemProp{}, fmt.Errorf("failed to get props : %w", err)
	}
	return props, nil
}
