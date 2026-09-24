package cart_transport_http

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type CartHTTPHandler struct {
	cartService CartService
}

type CartService interface {
	CalcItemFromCart(ctx context.Context, input model.Cart) (model.Cart, error)
	DeleteFromCart(ctx context.Context, input model.Cart) ([]model.Cart, error)
	AddToCart(ctx context.Context, input model.Cart) (model.Cart, error)
	GetCart(ctx context.Context, id string) ([]model.Cart, error)
}

func NewCartHTTPHandler(cartService CartService) *CartHTTPHandler {
	return &CartHTTPHandler{
		cartService: cartService,
	}
}
