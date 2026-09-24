package cart_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *CartService) DeleteFromCart(
	ctx context.Context,
	input model.Cart) ([]model.Cart, error) {

	cart, err := s.cartRepository.DeleteFromCart(ctx, input)
	if err != nil {
		return []model.Cart{}, fmt.Errorf("failed to delete item from cart id ='%d', item id ='%d' : %w", input.Id, input.ItemId, err)
	}
	return cart, nil
}
