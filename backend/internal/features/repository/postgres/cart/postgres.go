package cart_repository_postgres

import core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"

type CartRepository struct {
	pool core_postgres_pool.Pool
}

func NewCartRepository(pool core_postgres_pool.Pool) *CartRepository {
	return &CartRepository{
		pool: pool,
	}
}
