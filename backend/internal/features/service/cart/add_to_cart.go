package cart_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *CartService) AddToCart(
	ctx context.Context,
	input model.Cart) (model.Cart, error) {

	cart, err := s.cartRepository.AddToCart(ctx, input)
	if err != nil {
		return model.Cart{}, fmt.Errorf("failed addToCart id='%d' :%w", input.Id, err)
	}

	return cart, nil

}
