package item_transport_http

import (
	"context"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

type ItemHTTPHandler struct {
	itemService ItemService
}

type ItemService interface {
	CreateItem(
		ctx context.Context,
		input model.Item,
		documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error)

	GetById(
		ctx context.Context,
		id int) (model.ItemProp, error)

	GetSpecById(
		ctx context.Context,
		id int) ([]model.ItemSpecFiles, error)

	GetItemId(
		ctx context.Context,
		input model.Item) (model.Item, error)

	DeleteItem(
		ctx context.Context,
		input model.Item) (model.Item, error)

	GetAll(
		ctx context.Context) ([]model.Item, error)

	ChangeVisible(
		ctx context.Context,
		input bool,
		id string) (model.Item, error)

	GetAllPropsName(
		ctx context.Context) (model.ItemProp, error)

	SetProps(
		ctx context.Context,
		input []model.ItemProp,
		id string) ([]model.ItemProp, error)
}

func NewItemHTTPHandler(itemService ItemService) *ItemHTTPHandler {
	return &ItemHTTPHandler{
		itemService: itemService,
	}
}
