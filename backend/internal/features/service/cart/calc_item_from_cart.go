package cart_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *CartService) CalcItemFromCart(
	ctx context.Context,
	input model.Cart) (model.Cart, error) {

	cart, err := s.cartRepository.CalcItemFromCart(ctx, input)
	if err != nil {
		return model.Cart{}, fmt.Errorf("failed to calc item cart id='%d', item id ='%d' :%w", input.Id, input.ItemId, err)
	}
	return cart, nil
}
