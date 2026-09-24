package cart_service

import (
	"context"
	"fmt"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (s *CartService) GetCart(
	ctx context.Context,
	id string) ([]model.Cart, error) {

	cart, err := s.cartRepository.GetCart(ctx, id)
	if err != nil {
		return []model.Cart{}, fmt.Errorf("failed to get cart id ='%d' :%w", id, err)
	}
	return cart, nil
}
