package usecase

import (
	"myWebApp/backend/internal/domain/model"
)

type ItemUseCase struct {
	repo model.ItemRepository
}

func NewItemUseCase(r model.ItemRepository) *ItemUseCase {
	return &ItemUseCase{repo: r}
}

func (i *ItemUseCase) CreateItem(input model.Item, documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error) {
	return i.repo.CreateItem(input, documents)
}

func (i *ItemUseCase) GetItemById(id int) (model.ItemProp, error) {
	return i.repo.GetItemById(id)
}

func (i *ItemUseCase) GetSpecById(id int) ([]model.ItemSpecFiles, error) {
	return i.repo.GetSpecById(id)
}
func (i *ItemUseCase) GetItemId(input model.Item) (model.Item, error) {
	return i.repo.GetItemId(input)
}
func (i *ItemUseCase) DeleteItem(input model.Item) (model.Item, error) {
	return i.repo.DeleteItem(input)
}
func (i *ItemUseCase) GetAllItems() ([]model.Item, error) {
	return i.repo.GetAllItems()
}
func (i *ItemUseCase) ChangeVisible(input bool, id string) (model.Item, error) {
	return i.repo.ChangeVisible(input, id)
}
func (i *ItemUseCase) GetAllPropsName() (model.ItemProp, error) {
	return i.repo.GetAllPropsName()
}

func (i *ItemUseCase) SetPropsForItem(input []model.ItemProp, id string) ([]model.ItemProp, error) {
	return i.repo.SetPropsForItem(input, id)
}
