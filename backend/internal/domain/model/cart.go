package model

import "context"

type Cart struct {
	Id                int            `json:"id"`
	ItemSpecId        int            `json:"item_spec_id"`
	ItemId            int            `json:"item_id"`
	Name              string         `json:"name"`
	Price             int            `json:"price"`
	ItemType          string         `json:"item_type"`
	SecondaryItemType string         `json:"item_secondary_type"`
	Count             int            `json:"count"`
	ItemPicture       string         `json:"item_picture"`
	ItemDescription   string         `json:"item_description"`
	Article           string         `json:"item_article"`
	Props             map[string]any `json:"props"`
}

type CartRepository interface {
	AddItemToCart(ctx context.Context, input Cart) (Cart, error)
	DeleteShopCart(ctx context.Context, input string) (Cart, error)
	GetCart(ctx context.Context, id string) ([]Cart, error)
	DeleteUserItem(ctx context.Context, input Cart) ([]Cart, error)
	CalcUserItem(ctx context.Context, input Cart) (Cart, error)
}
