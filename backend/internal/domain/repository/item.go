package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"log/slog"
	"os"
)

type itemRepo struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewItemRepo(db *pgxpool.Pool, logger *slog.Logger) model.ItemRepository {
	return &itemRepo{
		db:     db,
		logger: logger,
	}
}

func (s *itemRepo) GetAllItems(ctx context.Context) ([]model.Item, error) {
	var items []model.Item
	s.logger.Debug("Repository: GetAllItems(getting all goods from database) started")
	rows, err := s.db.Query(ctx, `SELECT *
       FROM item`)
	if err != nil {
		s.logger.Error("Error when trying to get all items from handler GetAllItems", slog.Any("error: ", err), slog.Any("db_error: ", err))
		return items, err
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&item.ItemType,
			&item.ItemSecondaryType,
			&item.ItemPicture,
			&item.ItemDescription,
			&item.ItemShortDescription,
			&item.Article,
			&item.Visible)
		if err != nil {
			s.logger.Error("Error when scanning all items to a struct from handler GetAllItems", slog.Any("error: ", err), slog.Any("db_error: ", err))
			return items, err
		}
		items = append(items, item)
	}
	s.logger.Debug("Repository: GetAllItems success")
	return items, nil
}

func (s *itemRepo) DeleteItem(ctx context.Context, input model.Item) (model.Item, error) {
	s.logger.Debug("Repository: DeleteItem(delete item from database) started", "item_id: ", input.Id)
	var item model.Item
	var document model.ItemSpecFiles
	err := s.db.QueryRow(ctx, `SELECT link FROM item_spec_files WHERE id = $1`, input.Id).Scan(&document.SpecFileLink)
	if err != nil {
		s.logger.Error("Repository: DeleteItem(delete item from database) error when search link(specification) for item", slog.Any("item_id: ", input.Id), slog.Any("db_error: ", err))
		return item, err
	}
	err = os.Remove("./public/documents/" + document.SpecFileLink)
	if err != nil {
		s.logger.Error("Repository: DeleteItem(delete item from database) error when delete specification file from /public/documents", slog.Any("item_id: ", input.Id), slog.Any("db_error: ", err))
		return item, err
	}

	err = s.db.QueryRow(ctx, `SELECT item_picture FROM item WHERE id =$1`, input.Id).Scan(&item.ItemPicture)
	if err != nil {
		s.logger.Error("Repository: DeleteItem(delete item from database) error when search item picture link", slog.Any("item_id: ", input.Id), slog.Any("db_error: ", err))
		return item, err
	}
	err = os.Remove("./public/images/" + item.ItemPicture)
	if err != nil {
		s.logger.Error("Repository: DeleteItem(delete item from database) error when delete item picture", slog.Any("item_id: ", input.Id), slog.Any("db_error: ", err))
		return item, err
	}
	_, err = s.db.Exec(ctx, `DELETE FROM item WHERE id =$1`, input.Id)
	if err != nil {
		s.logger.Error("Repository: DeleteItem(delete item from database) error when deleting item rest information", slog.Any("item_id: ", input.Id), slog.Any("db_error: ", err))
		return item, err
	}
	s.logger.Debug("Repository: DeleteItem(delete item from database) success", "item_id: ", input.Id)
	return item, nil
}

func (s *itemRepo) GetItemId(ctx context.Context, input model.Item) (model.Item, error) {
	var item model.Item
	s.logger.Debug("Repository: GetItemId (getting id from item name) working", "item_name: ", input.Name)

	err := s.db.QueryRow(ctx,
		"SELECT id FROM item WHERE name =$1", input.Name).Scan(&item.Id)

	if err != nil {
		s.logger.Error("Repository: GetItemId (getting id from item name) error when search item_id in db", slog.Any("item_name:", input.Name), slog.Any("db_error: ", err))
		return item, err
	}
	s.logger.Debug("Repository: GetItemId (getting id from item name) success", "item_name: ", input.Name)
	return item, nil
}

func (s *itemRepo) GetSpecById(ctx context.Context, id int) ([]model.ItemSpecFiles, error) {
	s.logger.Debug("Repository: GetSpecById (getting spec file link from db using item_id) working", "item_id: ", id)
	var specStore []model.ItemSpecFiles
	rows, err := s.db.Query(ctx, `SELECT 
    name, link, picture 
    FROM item_spec_files 
    WHERE id =$1`, id)
	if err != nil {
		s.logger.Error("Repository: GetSpecById (getting spec file link from db using item_id) error when search item on db", slog.Any("item_id:", id), slog.Any("db_error: ", err))
		return specStore, err
	}
	defer rows.Close()
	for rows.Next() {
		var specRow model.ItemSpecFiles
		err = rows.Scan(&specRow.SpecFileName,
			&specRow.SpecFileLink,
			&specRow.SpecFilePic)
		if err != nil {
			s.logger.Error("Repository: GetSpecById (getting spec file link from db using item_id) error when search item on db", slog.Any("item_id:", id), slog.Any("db_error: ", err))
			return specStore, err
		}
		specStore = append(specStore, specRow)
	}
	s.logger.Debug("Repository: GetSpecById (getting spec file link from db using item_id) success", "item_id: ", id)
	return specStore, nil
}

func (s *itemRepo) GetItemById(ctx context.Context, id int) (model.ItemProp, error) {
	s.logger.Debug("Repository: GetItemById (getting item using id) working", "id: ", id)

	var item model.ItemProp
	err := s.db.QueryRow(ctx, `SELECT 
    item.id,
    item.name, 
    item.price, 
    item.item_description, 
    item.item_short_description,
    item.item_picture,
    item.article,
    item.item_type,
    item.secondary_type,
    item.visible
    FROM item WHERE item.id =$1`, id).
		Scan(&item.Id,
			&item.Name,
			&item.Price,
			&item.ItemDescription,
			&item.ItemShortDescription,
			&item.ItemPicture,
			&item.Article,
			&item.ItemType,
			&item.ItemSecondaryType,
			&item.Visible)
	if err != nil {
		s.logger.Error("Repository: GetItemById (getting item using id)  error when search item on db", slog.Any("id:", id), slog.Any("db_error: ", err))
		return item, err
	}
	rows, err := s.db.Query(ctx,
		`SELECT 
    item_properties.name,
    item_properties_values.value 
    FROM item 
    JOIN item_properties_values ON item.id = item_properties_values.item_id 
        JOIN item_properties ON item_properties_values.property_id = item_properties.id 
    WHERE item.id =$1`, id)
	if err != nil {

		s.logger.Error("Repository: GetItemById (getting item using id)  error when search item props on db", slog.Any("id:", id), slog.Any("db_error: ", err))
		return item, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var prop string

		err = rows.Scan(&name, &prop)
		if err != nil {
			s.logger.Error("Repository: GetItemById (getting item using id)  error when scanning props from db", slog.Any("id:", id), slog.Any("db_error: ", err))
			return item, err
		}
		item.PropNameA = append(item.PropNameA, name)
		item.PropValueA = append(item.PropValueA, prop)
	}
	s.logger.Debug("Repository: GetItemById (getting item using id) success", "id: ", id)
	return item, nil
}

func (s *itemRepo) CreateItem(ctx context.Context, input model.Item, documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error) {
	var newItem model.Item
	var document model.ItemSpecFiles
	s.logger.Debug("Repository: CreateItem working")
	err := s.db.QueryRow(ctx,
		`INSERT INTO item (name,price,item_type, secondary_type, item_picture, item_description, item_short_description, article) 
        VALUES ($1,$2,$3,$4,$5,$6,$7, $8) RETURNING id,name,price,item_type`,
		input.Name, input.Price, input.ItemType, input.ItemSecondaryType, input.ItemPicture, input.ItemDescription, input.ItemShortDescription, input.Article).
		Scan(&newItem.Id, &newItem.Name, &newItem.Price, &newItem.ItemType)
	if err != nil {
		s.logger.Error("Repository: CreateItem error when insert item info", slog.Any("db_error: ", err))
		return newItem, document, err
	}

	_, err = s.db.Exec(ctx,
		`INSERT INTO item_spec_files (id,name,link,picture) VALUES ($1, $2, $3,$4)`,
		newItem.Id, documents.SpecFileName, documents.SpecFileLink, documents.SpecFilePic)
	if err != nil {
		s.logger.Error("Repository: CreateItem error when insert specification file for item", slog.Any("db_error: ", err))
		return newItem, document, err
	}
	s.logger.Debug("Repository: CreateItem success")
	return newItem, document, nil
}

func (s *itemRepo) ChangeVisible(ctx context.Context, status bool, id string) (model.Item, error) {
	var visible model.Item
	s.logger.Debug("Repository: ChangeVisible for item working", "item_id: ", id)
	err := s.db.QueryRow(ctx, `UPDATE item SET visible = $1 WHERE id = $2 RETURNING visible`, status, id).Scan(&visible.Visible)
	if err != nil {
		s.logger.Error("Repository:ChangeVisible error when search item", slog.Any("db_error: ", err))
		return visible, err
	}
	s.logger.Debug("Repository: ChangeVisible for item success", "item_id: ", id)
	return visible, nil
}

func (s *itemRepo) GetAllPropsName(ctx context.Context) (model.ItemProp, error) {
	var propsName model.ItemProp
	s.logger.Debug("Repository: GetAllPropsName for admin working")
	temp := make([]string, 30)
	var index int
	rows, err := s.db.Query(ctx, `SELECT * FROM item_properties`)
	if err != nil {
		s.logger.Error("Repository: GetAllPropsName error when get all item properties for Admin", slog.Any("db_error: ", err))
		return propsName, err
	}
	defer rows.Close()
	for rows.Next() {
		var prop string
		err = rows.Scan(&index, &prop)
		if err != nil {
			s.logger.Error("Repository: GetAllPropsName error when scan all item properties for Admin", slog.Any("db_error: ", err))
			return propsName, err
		}
		temp[index] = prop
	}

	propsName.PropNameA = temp[:index]
	s.logger.Debug("Repository: GetAllPropsName for admin success")
	return propsName, nil
}

func (s *itemRepo) SetPropsForItem(ctx context.Context, input []model.ItemProp, id string) ([]model.ItemProp, error) {
	var props []model.ItemProp
	s.logger.Debug("Repository: SetPropsForItem working", "item_id: ", id)
	idSlice := make([]int, 0, len(input))
	valueSlice := make([]string, 0, len(input))

	for _, item := range input {
		idSlice = append(idSlice, item.IdProp)
		valueSlice = append(valueSlice, item.PropValue)
	}

	_, err := s.db.Exec(ctx, `INSERT INTO item_properties_values (item_id,property_id, value) 
      SELECT $1, unnest($2::int[]), unnest($3::text[]) `,
		id, idSlice, valueSlice)
	if err != nil {
		s.logger.Error("Repository: SetPropsForItem error when insert new props", slog.Any("db_error: ", err))
		return input, err
	}
	s.logger.Debug("Repository: SetPropsForItem success", "item_id: ", id)
	return props, nil
}
