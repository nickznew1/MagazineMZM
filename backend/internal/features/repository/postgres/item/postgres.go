package item_repository_postgres

import core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"

type ItemRepository struct {
	pool core_postgres_pool.Pool
}

func NewItemRepository(pool core_postgres_pool.Pool) *ItemRepository {
	return &ItemRepository{
		pool: pool,
	}
}
