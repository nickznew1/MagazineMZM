package users_repository_postgres

import (
	"context"
	"log/slog"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (r *UsersRepository) GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error) {

	var users []model.UserOrdinaryInfo
	rows, err := r.db.Query(ctx, "SELECT id,login,password,email, registration_date,user_role FROM customer")
	if err != nil {
		r.logger.Error("Repository: GetAllUsers (for admin) error when select all user from customer table", slog.Any("db_err: ", err))
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserOrdinaryInfo
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegistrationDate, &user.UserRole)
		if err != nil {
			r.logger.Error("Repository: GetAllUsers (for admin) error when append all user to result slice", slog.Any("db_err: ", err))
			return users, err
		}
		users = append(users, user)
	}
	r.logger.Debug("Repository: GetAllUsers (for admin) success")
	return users, nil
}
