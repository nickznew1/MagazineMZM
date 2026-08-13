package usecase

import (
	"context"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type CartUseCase struct {
	repo model.CartRepository
}

func NewCartUseCase(r model.CartRepository) *CartUseCase {
	return &CartUseCase{repo: r}
}

func (c *CartUseCase) CalcUserItem(ctx context.Context, input model.Cart) (model.Cart, error) {
	return c.repo.CalcUserItem(ctx, input)
}

func (c *CartUseCase) DeleteUserItem(ctx context.Context, input model.Cart) ([]model.Cart, error) {
	return c.repo.DeleteUserItem(ctx, input)
}

func (c *CartUseCase) AddItemToCart(ctx context.Context, input model.Cart) (model.Cart, error) {
	return c.repo.AddItemToCart(ctx, input)
}

func (c *CartUseCase) DeleteShopCart(ctx context.Context, input string) (model.Cart, error) {
	return c.repo.DeleteShopCart(ctx, input)
}

func (c *CartUseCase) GetCart(ctx context.Context, id string) ([]model.Cart, error) {
	return c.repo.GetCart(ctx, id)
}
