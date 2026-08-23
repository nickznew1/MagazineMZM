package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

type userRepo struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewUserRepo(db *pgxpool.Pool, logger *slog.Logger) model.UserRepository {
	return &userRepo{
		logger: logger,
		db:     db}
}

func (r *userRepo) GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	r.logger.Debug("Repository: GetUserById (Getting user by login) started", "user_login: ", input.Login)
	err := r.db.QueryRow(ctx, "SELECT id,login, password FROM customer WHERE login =$1", input.Login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		r.logger.Error("Repository: GetUserById (Getting user by login) error", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: GetUserById (Getting user by login) success", "user_login: ", input.Login)
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var customer model.UserOrdinaryInfo
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	r.logger.Debug("Repository: CreateUser started")
	if err != nil {
		r.logger.Error("Repository: CreateUser error - can't hash password(bcrypt)", slog.Any("bcrypt_err: ", err))
		return customer, err
	}
	input.Password = string(hashedPassword)

	err = r.db.QueryRow(ctx,
		"INSERT INTO customer (login,password, email) VALUES ($1,$2,$3) RETURNING id,login,password,email",
		input.Login, input.Password, input.Email).
		Scan(&customer.Id, &customer.Login, &customer.Password, &customer.Email)
	if err != nil {
		r.logger.Error("Repository: CreateUser error when insert new user into customer table", slog.Any("db_err: ", err))
		return customer, err
	}

	r.logger.Debug("Repository: CreateUser success", "new user_id: ", customer.Id)
	return customer, nil
}

func (r *userRepo) UserAuth(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	r.logger.Debug("Repository: UserAuth started", "user_id: ", input.Id)

	user, err := r.GetUserById(ctx, input)
	if err != nil {
		r.logger.Error("Repository: UserAuth error - can t get user info when check user_id", slog.Any("db_err: ", err))
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		r.logger.Error("Repository: UserAuth error - wrong password for user", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: UserAuth success", "user_id: ", input.Id)
	return user, nil
}

func (r *userRepo) FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error) {
	var profileInfo model.UserOrdinaryInfoOut
	r.logger.Debug("Repository: FetchProfileInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,login,email,user_role from customer WHERE id =$1", id).Scan(&profileInfo.Id, &profileInfo.Login, &profileInfo.Email, &profileInfo.UserRole)
	if err != nil {
		r.logger.Error("Repository: FetchProfileInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profileInfo, err
	}
	r.logger.Debug("Repository: FetchProfileInfo success (goroutines) ", "user_id: ", id)
	return profileInfo, nil
}

func (r *userRepo) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	var profilePersonalInfo model.UserPersonalInfoOut
	r.logger.Debug("Repository: FetchProfilePersonalInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,company, first_name, second_name from customer_personal_info WHERE id=$1", id).Scan(&profilePersonalInfo.Id, &profilePersonalInfo.Company, &profilePersonalInfo.FirstName, &profilePersonalInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: FetchProfilePersonalInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profilePersonalInfo, err
	}
	r.logger.Debug("Repository: FetchProfilePersonalInfo success (goroutines)", "user_id: ", id)
	return profilePersonalInfo, nil
}

func (r *userRepo) FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error) {
	var profileDeliveryInfo model.UserDeliveryInfoOut
	r.logger.Debug("Repository: FetchProfileDeliveryInfo started (goroutines)", "user_id: ", id)
	err := r.db.QueryRow(ctx, "SELECT id,phone_number, city, address from customer_delivery_info WHERE id = $1", id).Scan(&profileDeliveryInfo.Id, &profileDeliveryInfo.PhoneNumber, &profileDeliveryInfo.City, &profileDeliveryInfo.Address)
	if err != nil {
		r.logger.Error("Repository: FetchProfileDeliveryInfo error - can t find user info with id", slog.Any("db_err: ", err))
		return profileDeliveryInfo, err
	}
	r.logger.Debug("Repository: FetchProfileDeliveryInfo success (goroutines)", "user_id: ", id)
	return profileDeliveryInfo, nil
}

func (r *userRepo) RecordPersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	r.logger.Debug("Repository: RecordPersonalInfo (record new personal info for user) started", "input :", input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_personal_info (id,company, first_name, second_name) VALUES ($1, $2, $3, $4) RETURNING id, company, first_name,second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Id, &userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when insert new personal info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: RecordPersonalInfo (record new personal info for user) success", "input :", input)
	return userInfo, nil
}

func (r *userRepo) UpdatePersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	r.logger.Debug("Repository: UpdatePersonalInfo (update personal info for user) started", "input :", input)

	err := r.db.QueryRow(ctx, "UPDATE customer_personal_info SET company =$2, first_name=$3, second_name=$4 WHERE id =$1 RETURNING company, first_name, second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when update personal info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: UpdatePersonalInfo (update personal info for user) success", "input :", input)
	return userInfo, nil
}

func (r *userRepo) RecordDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	r.logger.Debug("Repository: RecordDeliveryInfo started", "input :", input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_delivery_info (id,phone_number, city, address) VALUES ($1, $2, $3, $4) RETURNING id, phone_number, city,address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.Id, &userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when record delivery info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: RecordDeliveryInfo success", "input :", input)
	return userInfo, nil
}

func (r *userRepo) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	r.logger.Debug("Repository: UpdateDeliveryInfo started", "input :", input)
	err := r.db.QueryRow(ctx, "UPDATE customer_delivery_info SET phone_number=$2, city=$3, address=$4 WHERE id =$1 RETURNING phone_number, city, address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		r.logger.Error("Repository: RecordPersonalInfo error when update delivery info for user", slog.Any("db_err: ", err))
		return userInfo, err
	}
	r.logger.Debug("Repository: UpdateDeliveryInfo success", "input :", input)
	return userInfo, nil
}

func (r *userRepo) UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error) {
	var check model.PasswordChange
	r.logger.Debug("Repository: UserPasswordChange started", "input :", input)
	err := r.db.QueryRow(ctx, "SELECT password FROM customer WHERE id =$1", input.Id).Scan(&check.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(check.OldPassword), []byte(input.OldPassword))
	if err != nil {
		r.logger.Error("Repository: UserPasswordChange bcrypt error - user writes wrong ordinary password", slog.Any("bcrypt_err: ", err))
		return check, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 12)
	err = r.db.QueryRow(ctx, "UPDATE customer SET password = $1 WHERE id = $2 RETURNING password,id", hashedPassword, input.Id).Scan(&check.NewPassword, &check.Id)
	if err != nil {
		r.logger.Error("Repository: UserPasswordChange bcrypt error when update password for user", slog.Any("db_err: ", err))
		return check, err
	}
	r.logger.Debug("Repository: UserPasswordChange success", "input :", input)
	return check, nil
}

func (r *userRepo) UserChangeEmail(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	var loginExists bool
	r.logger.Debug("Repository: UserChangeEmail started", "input:", input)
	err := r.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM customer WHERE login =$1)", input.Login).Scan(&loginExists)
	if err != nil {
		r.logger.Error("Repository: UserChangeEmail error when update email for user - login doesn't found", slog.Any("db_err: ", err))
		return user, err
	}
	if !loginExists {
		r.logger.Error("Repository: UserChangeEmail error when update email for user - login doesn't found", slog.Any("db_err: ", err))
		return user, err
	}
	err = r.db.QueryRow(ctx, "UPDATE customer SET email = $1 WHERE login = $2 RETURNING login,email", input.Email, input.Login).Scan(&user.Login, &user.Email)
	if err != nil {
		r.logger.Error("Repository: UserChangeEmail error when update email for user (db error)", slog.Any("db_err: ", err))
		return user, err
	}
	r.logger.Debug("Repository: UserChangeEmail success", "input:", input)
	return user, nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error) {
	r.logger.Debug("Repository: GetAllUsers (for admin) started")
	var users []model.UserOrdinaryInfo
	rows, err := r.db.Query(ctx, "SELECT id,login,password,email, registration_date,user_role FROM customer")
	if err != nil {
		r.logger.Error("Repository: GetAllUsers (for admin) error when select all users from customer table", slog.Any("db_err: ", err))
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserOrdinaryInfo
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegistrationDate, &user.UserRole)
		if err != nil {
			r.logger.Error("Repository: GetAllUsers (for admin) error when append all users to result slice", slog.Any("db_err: ", err))
			return users, err
		}
		users = append(users, user)
	}
	r.logger.Debug("Repository: GetAllUsers (for admin) success")
	return users, nil
}
