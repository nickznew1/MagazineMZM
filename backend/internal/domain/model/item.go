package model

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
	CreateItem(input Item, documents ItemSpecFiles) (Item, ItemSpecFiles, error)
	GetItemById(id int) (ItemProp, error)
	GetSpecById(id int) ([]ItemSpecFiles, error)
	GetItemId(input Item) (Item, error)
	DeleteItem(input Item) (Item, error)
	GetAllItems() ([]Item, error)
	ChangeVisible(status bool, id string) (Item, error)
	GetAllPropsName() (ItemProp, error)
	SetPropsForItem(input []ItemProp, id string) ([]ItemProp, error)
}
