package repository

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"strconv"
)

type cartRepo struct {
	logger *slog.Logger
	db     *pgxpool.Pool
}

func NewCartRepo(db *pgxpool.Pool, logger *slog.Logger) model.CartRepository {
	return &cartRepo{
		logger: logger,
		db:     db}
}

func (s *cartRepo) CalcUserItem(ctx context.Context, input model.Cart) (model.Cart, error) {
	var calc model.Cart
	s.logger.Debug("Repository: CalcUserItem (increase/decrease item count in user cart) started", "input: ", input)
	err := s.db.QueryRow(ctx, `UPDATE customer_item 
  SET count =$1 
  WHERE item_id =$2 
  AND customer_id = $3 
  RETURNING item_id,count, customer_id`,
		input.Count, input.ItemId, input.Id).Scan(&calc.ItemSpecId, &calc.Count, &calc.Id)
	if err != nil {
		s.logger.Error("Repository: CalcUserItem error when update customer_item table", slog.Any("db_error: ", err))
		return calc, err
	}
	s.logger.Debug("Repository: CalcUserItem (increase/decrease item count in user cart) success", "input: ", input)
	return calc, err

}

func (s *cartRepo) DeleteUserItem(ctx context.Context, input model.Cart) ([]model.Cart, error) {
	var cart []model.Cart
	s.logger.Debug("Repository: DeleteUserItem (deleting item from user cart) started", "input: ", input)
	_, err := s.db.Exec(ctx, "DELETE FROM customer_item WHERE item_id = $1 AND customer_id = $2 ", input.ItemId, input.Id)
	if err != nil {
		s.logger.Error("Repository: DeleteUserItem (deleting item from user cart) error when delete from customer_item", slog.Any("db_error: ", err))
		return cart, err
	}
	newCart, err := s.GetCart(ctx, strconv.Itoa(input.Id))
	if err != nil {
		s.logger.Error("Repository: DeleteUserItem error when trying to get new cart store after deleting", slog.Any("db_error: ", err))
	}
	s.logger.Debug("Repository: DeleteUserItem (deleting item from user cart) success", "input: ", input)
	return newCart, err
}

func (s *cartRepo) AddItemToCart(ctx context.Context, input model.Cart) (model.Cart, error) {
	var userItem model.Cart
	s.logger.Debug("Repository: AddItemToCart (user added new item to his cart) started", "input: ", input)
	propsJson, _ := json.Marshal(input.Props)
	err := s.db.QueryRow(ctx, `
        INSERT INTO customer_item(customer_id, count, item_id, props) 
        VALUES($1, $2, $3, $4) 
        RETURNING customer_id, count, item_id
    `, input.Id, input.Count, input.ItemId, propsJson).Scan(
		&userItem.Id, &userItem.Count, &userItem.ItemId)
	if err != nil {
		s.logger.Error("Repository: ddItemToCart (user added new item to his cart) error when trying to insert new item to cart", slog.Any("db_error: ", err))
		return userItem, err
	}
	s.logger.Debug("Repository: AddItemToCart (user added new item to his cart) success", "input: ", input)
	return userItem, nil
}

func (s *cartRepo) DeleteShopCart(ctx context.Context, input string) (model.Cart, error) {
	var userItem model.Cart
	s.logger.Debug("Repository: DeleteShopCart (for admin) started", "input: ", input)
	err := s.db.QueryRow(ctx, "DELETE FROM customer_item WHERE login =$1", input)
	if err != nil {
		s.logger.Error("Repository: DeleteShopCart (for admin) started error when trying to delete shopping cart", slog.Any("db_error: ", err))
		return userItem, nil
	}
	s.logger.Debug("Repository: DeleteShopCart (for admin) success", "input: ", input)
	return userItem, nil
}

func (s *cartRepo) GetCart(ctx context.Context, id string) ([]model.Cart, error) {
	var shoppingCart []model.Cart
	s.logger.Debug("Repository: GetCart started", "user_id : ", id)
	rows, err := s.db.Query(ctx, `SELECT
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
		s.logger.Error("Repository: GetCart error when trying to get shopping cart for user", slog.Any("db_error: ", err))
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
			s.logger.Error("Repository: GetCart error when trying to append items to result slice", slog.Any("db_error: ", err))
			return shoppingCart, err
		}
		err = json.Unmarshal(raw, &props)
		cart.Props = props
		shoppingCart = append(shoppingCart, cart)
	}
	s.logger.Debug("Repository: GetCart success", "user_id : ", id)
	return shoppingCart, nil
}
