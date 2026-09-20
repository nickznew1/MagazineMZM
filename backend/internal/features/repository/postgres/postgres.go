package users_repository_postgres

type UsersRepository struct {
	pool
}

func NewUsersRepository(pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
