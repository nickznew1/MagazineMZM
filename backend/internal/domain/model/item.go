package model

import "context"

type Item struct {
	Id                   int      `json:"id"`
	Name                 string   `json:"name"`
	Price                string   `json:"price"`
	ItemType             string   `json:"item_type"`
	ItemSecondaryType    string   `json:"item_secondary_type"`
	ItemPicture          string   `json:"item_picture"`
	ItemDescription      string   `json:"item_description"`
	ItemShortDescription string   `json:"item_short_description"`
	Clicks               int      `json:"clicks"`
	Article              string   `json:"article"`
	Visible              bool     `json:"visible"`
	SpecFileName         []string `json:"spec_file_name"`
	SpecFileLink         []string `json:"spec_file_link"`
	SpecFilePic          []string `json:"spec_file_pic"`
}

type ItemSpecFiles struct {
	SpecFileName string `json:"spec_file_name"`
	SpecFileLink string `json:"spec_file_link"`
	SpecFilePic  string `json:"spec_file_pic"`
}

type ItemProp struct {
	Item
	IdProp     int      `json:"prop_id"`
	PropValue  string   `json:"prop_value"`
	PropNameA  []string `json:"prop_name_array"`
	PropValueA []string `json:"prop_value_array"`
}

type ItemRepository interface {
	CreateItem(ctx context.Context, input Item, documents ItemSpecFiles) (Item, ItemSpecFiles, error)
	GetItemById(ctx context.Context, id int) (ItemProp, error)
	GetSpecById(ctx context.Context, id int) ([]ItemSpecFiles, error)
	GetItemId(ctx context.Context, input Item) (Item, error)
	DeleteItem(ctx context.Context, input Item) (Item, error)
	GetAllItems(ctx context.Context) ([]Item, error)
	ChangeVisible(ctx context.Context, status bool, id string) (Item, error)
	GetAllPropsName(ctx context.Context) (ItemProp, error)
	SetPropsForItem(ctx context.Context, input []ItemProp, id string) ([]ItemProp, error)
}
