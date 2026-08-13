package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"os"
)

type itemRepo struct {
	db *pgxpool.Pool
}

func NewItemRepo(db *pgxpool.Pool) model.ItemRepository {
	return &itemRepo{
		db: db}
}

func (s *itemRepo) GetAllItems(ctx context.Context) ([]model.Item, error) {
	var items []model.Item
	fmt.Println("get all items")
	rows, err := s.db.Query(ctx, `SELECT *
       FROM item`)
	if err != nil {
		fmt.Println(err)
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
			fmt.Println(err)
			return items, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *itemRepo) DeleteItem(ctx context.Context, input model.Item) (model.Item, error) {
	fmt.Println("DeleteItem")
	var item model.Item
	var document model.ItemSpecFiles
	err := s.db.QueryRow(ctx, `SELECT link FROM item_spec_files WHERE id = $1`, input.Id).Scan(&document.SpecFileLink)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	err = os.Remove("./public/documents/" + document.SpecFileLink)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	fmt.Println("документ удален")
	err = s.db.QueryRow(ctx, `SELECT item_picture FROM item WHERE id =$1`, input.Id).Scan(&item.ItemPicture)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	err = os.Remove("./public/images/" + item.ItemPicture)
	if err != nil {
		fmt.Println(err)
		return item, err
	}
	_, err = s.db.Exec(ctx, `DELETE FROM item WHERE id =$1`, input.Id)
	if err != nil {
		fmt.Println(err)
		return item, err
	}

	return item, nil
}

func (s *itemRepo) GetItemId(ctx context.Context, input model.Item) (model.Item, error) {
	var item model.Item
	fmt.Println("Получение ID товара")
	fmt.Println(input.Name)
	err := s.db.QueryRow(ctx,
		"SELECT id FROM item WHERE name =$1", input.Name).Scan(&item.Id)

	if err != nil {
		fmt.Println("Oshibka")
		return item, err
	}
	return item, nil
}

func (s *itemRepo) GetSpecById(ctx context.Context, id int) ([]model.ItemSpecFiles, error) {
	fmt.Println("get specFiles for item", id)
	var specStore []model.ItemSpecFiles
	rows, err := s.db.Query(ctx, `SELECT 
    name, link, picture 
    FROM item_spec_files 
    WHERE id =$1`, id)
	if err != nil {
		fmt.Println("cant find spec files for item ", id)
		return specStore, err
	}
	defer rows.Close()
	for rows.Next() {
		var specRow model.ItemSpecFiles
		err = rows.Scan(&specRow.SpecFileName,
			&specRow.SpecFileLink,
			&specRow.SpecFilePic)
		if err != nil {
			fmt.Println(err)
			return specStore, err
		}
		specStore = append(specStore, specRow)
	}
	return specStore, nil
}

func (s *itemRepo) GetItemById(ctx context.Context, id int) (model.ItemProp, error) {
	fmt.Println("getting item id for QueryParam")
	fmt.Println(id)
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
		fmt.Println(err)
		fmt.Println("item not found")
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
		fmt.Println(err)
		fmt.Println("item not found")
		return item, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var prop string

		err = rows.Scan(&name, &prop)
		if err != nil {
			fmt.Println(err)
			return item, err
		}
		item.PropNameA = append(item.PropNameA, name)
		item.PropValueA = append(item.PropValueA, prop)
	}
	return item, nil
}

func (s *itemRepo) CreateItem(ctx context.Context, input model.Item, documents model.ItemSpecFiles) (model.Item, model.ItemSpecFiles, error) {
	var newItem model.Item
	var document model.ItemSpecFiles
	fmt.Println("Creating new Item")
	err := s.db.QueryRow(ctx,
		`INSERT INTO item (name,price,item_type, secondary_type, item_picture, item_description, item_short_description, article) 
        VALUES ($1,$2,$3,$4,$5,$6,$7, $8) RETURNING id,name,price,item_type`,
		input.Name, input.Price, input.ItemType, input.ItemSecondaryType, input.ItemPicture, input.ItemDescription, input.ItemShortDescription, input.Article).
		Scan(&newItem.Id, &newItem.Name, &newItem.Price, &newItem.ItemType)
	if err != nil {
		fmt.Println(err)
		return newItem, document, err
	}

	_, err = s.db.Exec(ctx,
		`INSERT INTO item_spec_files (id,name,link,picture) VALUES ($1, $2, $3,$4)`,
		newItem.Id, documents.SpecFileName, documents.SpecFileLink, documents.SpecFilePic)
	if err != nil {
		fmt.Println(err)
		return newItem, document, err
	}

	return newItem, document, nil
}

func (s *itemRepo) ChangeVisible(ctx context.Context, status bool, id string) (model.Item, error) {
	var visible model.Item
	fmt.Println(status, id)
	err := s.db.QueryRow(ctx, `UPDATE item SET visible = $1 WHERE id = $2 RETURNING visible`, status, id).Scan(&visible.Visible)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error when trying change visible of item")
		return visible, err
	}
	return visible, nil
}

func (s *itemRepo) GetAllPropsName(ctx context.Context) (model.ItemProp, error) {
	var propsName model.ItemProp
	temp := make([]string, 30)
	var index int
	rows, err := s.db.Query(ctx, `SELECT * FROM item_properties`)
	if err != nil {
		fmt.Println(err)
		return propsName, err
	}
	defer rows.Close()
	for rows.Next() {
		var prop string
		err = rows.Scan(&index, &prop)
		if err != nil {
			fmt.Println(err)
			return propsName, err
		}
		temp[index] = prop
	}

	propsName.PropNameA = temp[:index]

	return propsName, nil
}

func (s *itemRepo) SetPropsForItem(ctx context.Context, input []model.ItemProp, id string) ([]model.ItemProp, error) {
	var props []model.ItemProp

	idSlice := make([]int, 0, len(input))
	valueSlice := make([]string, 0, len(input))

	for _, item := range input {
		idSlice = append(idSlice, item.IdProp)
		valueSlice = append(valueSlice, item.PropValue)
	}
	fmt.Println(idSlice, "id")
	fmt.Println(valueSlice, "value")
	fmt.Println(id)

	_, err := s.db.Exec(ctx, `INSERT INTO item_properties_values (item_id,property_id, value) 
      SELECT $1, unnest($2::int[]), unnest($3::text[]) `,
		id, idSlice, valueSlice)
	if err != nil {
		fmt.Println(err)
		return input, err
	}
	return props, nil
}
