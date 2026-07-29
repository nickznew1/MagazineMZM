package usecase

import "myWebApp/backend/internal/domain/model"

type CartUseCase struct {
	repo model.CartRepository
}

func NewCartUseCase(r model.CartRepository) *CartUseCase {
	return &CartUseCase{repo: r}
}

func (c *CartUseCase) CalcUserItem(input model.Cart) (model.Cart, error) {
	return c.repo.CalcUserItem(input)
}

func (c *CartUseCase) DeleteUserItem(input model.Cart) ([]model.Cart, error) {
	return c.repo.DeleteUserItem(input)
}

func (c *CartUseCase) AddItemToCart(input model.Cart) (model.Cart, error) {
	return c.repo.AddItemToCart(input)
}

func (c *CartUseCase) DeleteShopCart(input string) (model.Cart, error) {
	return c.repo.DeleteShopCart(input)
}

func (c *CartUseCase) GetCart(id string) ([]model.Cart, error) {
	return c.repo.GetCart(id)
}
