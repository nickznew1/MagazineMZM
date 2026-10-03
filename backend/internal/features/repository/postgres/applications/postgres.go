package applications_repository_postgres

import core_postgres_pool "github.com/nickznew1/MagazineMZM/backend/internal/core/repository/postgres/pool"

type ApplicationRepository struct {
	pool core_postgres_pool.Pool
}

func NewApplicationRepository(
	pool core_postgres_pool.Pool) *ApplicationRepository {

	return &ApplicationRepository{
		pool: pool,
	}
}
