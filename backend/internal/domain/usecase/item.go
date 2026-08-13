package usecase

import (
	"context"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type ItemUseCase struct {
	repo model.ItemRepository
}

func NewItemUseCase(r model.ItemRepository) *ItemUseCase {
	return &ItemUseCase{repo: r}
}

func (i *ItemUseCase) CreateItem(ctx context.Context, input model.Item, documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error) {
	return i.repo.CreateItem(ctx, input, documents)
}

func (i *ItemUseCase) GetItemById(ctx context.Context, id int) (model.ItemProp, error) {
	return i.repo.GetItemById(ctx, id)
}

func (i *ItemUseCase) GetSpecById(ctx context.Context, id int) ([]model.ItemSpecFiles, error) {
	return i.repo.GetSpecById(ctx, id)
}
func (i *ItemUseCase) GetItemId(ctx context.Context, input model.Item) (model.Item, error) {
	return i.repo.GetItemId(ctx, input)
}
func (i *ItemUseCase) DeleteItem(ctx context.Context, input model.Item) (model.Item, error) {
	return i.repo.DeleteItem(ctx, input)
}
func (i *ItemUseCase) GetAllItems(ctx context.Context) ([]model.Item, error) {
	return i.repo.GetAllItems(ctx)
}
func (i *ItemUseCase) ChangeVisible(ctx context.Context, input bool, id string) (model.Item, error) {
	return i.repo.ChangeVisible(ctx, input, id)
}
func (i *ItemUseCase) GetAllPropsName(ctx context.Context) (model.ItemProp, error) {
	return i.repo.GetAllPropsName(ctx)
}

func (i *ItemUseCase) SetPropsForItem(ctx context.Context, input []model.ItemProp, id string) ([]model.ItemProp, error) {
	return i.repo.SetPropsForItem(ctx, input, id)
}
