package model

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
	AddItemToCart(input Cart) (Cart, error)
	DeleteShopCart(input string) (Cart, error)
	GetCart(id string) ([]Cart, error)
	DeleteUserItem(input Cart) ([]Cart, error)
	CalcUserItem(input Cart) (Cart, error)
}
