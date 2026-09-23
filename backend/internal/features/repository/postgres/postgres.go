package users_repository_postgres

import "github.com/jackc/pgx/v5/pgxpool"

type UsersRepository struct {
	pool *pgxpool.Pool
}

func NewUsersRepository(pool *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
