package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myWebApp/backend/internal/domain/model"
	"strconv"
)

type cartRepo struct {
	db *sql.DB
}

func NewCartRepo(db *sql.DB) model.CartRepository {
	return &cartRepo{
		db: db}
}

func (s *cartRepo) CalcUserItem(input model.Cart) (model.Cart, error) {
	var calc model.Cart

	err := s.db.QueryRow(`UPDATE customer_item 
  SET count =$1 
  WHERE item_spec_id =$2 
  AND customer_id = $3 
  RETURNING item_spec_id,count, customer_id`,
		input.Count, input.ItemSpecId, input.Id).Scan(&calc.ItemSpecId, &calc.Count, &calc.Id)
	if err != nil {
		fmt.Println("item id doesnt founded")
		return calc, err
	}
	return calc, err

}

func (s *cartRepo) DeleteUserItem(input model.Cart) ([]model.Cart, error) {
	var cart []model.Cart
	fmt.Println("delete UserItem id", input.ItemId, input.Id)
	_, err := s.db.Exec("DELETE FROM customer_item WHERE item_spec_id = $1 AND customer_id = $2 ", input.ItemSpecId, input.Id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("User/item doesn't founded")
		return cart, err
	}
	newCart, err := s.GetCart(strconv.Itoa(input.Id))
	if err != nil {
		fmt.Println("cart empty")
		fmt.Println(newCart)
	}
	fmt.Println(newCart)
	return newCart, err
}

func (s *cartRepo) AddItemToCart(input model.Cart) (model.Cart, error) {
	var userItem model.Cart
	fmt.Println("Add item to cart for user")
	fmt.Println("add new item")
	propsJson, _ := json.Marshal(input.Props)
	err := s.db.QueryRow(`
        INSERT INTO customer_item(customer_id, count, item_id, props) 
        VALUES($1, $2, $3, $4) 
        RETURNING customer_id, count, item_id
    `, input.Id, input.Count, input.ItemId, propsJson).Scan(
		&userItem.Id, &userItem.Count, &userItem.ItemId)
	if err != nil {
		fmt.Println(err)
		fmt.Println("oshibka dobavleniya tovara v cart")
		return userItem, err
	}

	return userItem, nil
}

func (s *cartRepo) DeleteShopCart(input string) (model.Cart, error) {
	var userItem model.Cart
	fmt.Println("Deleting shopcart")
	fmt.Println(input)
	err := s.db.QueryRow("DELETE FROM customer_item WHERE login =$1", input)
	if err != nil {
		return userItem, nil
	}
	return userItem, nil
}

func (s *cartRepo) GetCart(id string) ([]model.Cart, error) {
	var shoppingCart []model.Cart
	fmt.Println("Creating shopping cart for user: ", id)
	rows, err := s.db.Query(`SELECT
    item_id,
    count,
    customer_id, 
    props,
    item.name, 
    item.price, 
    item.item_picture, 
    item.item_description,
    item.article,
    item.item_type,
    item.secondary_type
    FROM customer_item 
        JOIN item ON customer_item.item_id = item.id 
        WHERE customer_id =$1`, id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("cart is empty")
		return shoppingCart, err
	}
	for rows.Next() {
		var cart model.Cart
		var raw []byte
		var props map[string]any
		err = rows.Scan(
			&cart.ItemId,
			&cart.Count,
			&cart.Id,
			&raw,
			&cart.Name,
			&cart.Price,
			&cart.ItemPicture,
			&cart.ItemDescription,
			&cart.Article,
			&cart.ItemType,
			&cart.SecondaryItemType)
		if err != nil {
			fmt.Println(err)
			return shoppingCart, err
		}
		err = json.Unmarshal(raw, &props)
		cart.Props = props
		shoppingCart = append(shoppingCart, cart)
	}
	return shoppingCart, nil
}
