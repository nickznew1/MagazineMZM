package cart_service

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type CartService struct {
	cartRepository CartRepository
}

type CartRepository interface {
	CalcItemFromCart(ctx context.Context, input model.Cart) (model.Cart, error)
	DeleteFromCart(ctx context.Context, input model.Cart) ([]model.Cart, error)
	AddToCart(ctx context.Context, input model.Cart) (model.Cart, error)
	GetCart(ctx context.Context, id string) ([]model.Cart, error)
}

func NewCartService(cartRepository CartRepository) *CartService {
	return &CartService{
		cartRepository: cartRepository,
	}
}
