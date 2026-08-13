package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"

	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) model.UserRepository {
	return &userRepo{
		db: db}
}

func (r *userRepo) GetUserById(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	fmt.Println("Getting user by ID")
	err := r.db.QueryRow(ctx, "SELECT id,login, password FROM customer WHERE login =$1", input.Login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		fmt.Println("user not found")
		return user, err
	}
	fmt.Println("get user with login ", user.Login)
	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var customer model.UserOrdinaryInfo
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		fmt.Println("cant hash password")
		return customer, err
	}
	input.Password = string(hashedPassword)

	err = r.db.QueryRow(ctx,
		"INSERT INTO customer (login,password, email) VALUES ($1,$2,$3) RETURNING id,login,password,email",
		input.Login, input.Password, input.Email).
		Scan(&customer.Id, &customer.Login, &customer.Password, &customer.Email)
	if err != nil {
		fmt.Println(err)
		return customer, err
	}

	fmt.Println("user created", input.Login)
	return customer, nil
}

func (r *userRepo) UserAuth(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	fmt.Println("user authentication")
	fmt.Println(input)
	user, err := r.GetUserById(input)
	if err != nil {
		fmt.Println("user not found")
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		fmt.Println(err)
		fmt.Println("wrong password")
		return user, err
	}
	fmt.Println("user Auth", input.Id)
	return user, nil
}

func (r *userRepo) FetchProfileInfo(ctx context.Context, id string) (model.UserOrdinaryInfoOut, error) {
	var profileInfo model.UserOrdinaryInfoOut
	err := r.db.QueryRow(ctx, "SELECT id,login,email,user_role from customer WHERE id =$1", id).Scan(&profileInfo.Id, &profileInfo.Login, &profileInfo.Email, &profileInfo.UserRole)
	if err != nil {
		fmt.Println("user doesnt found")
		return profileInfo, err
	}
	return profileInfo, nil
}

func (r *userRepo) FetchProfilePersonalInfo(ctx context.Context, id string) (model.UserPersonalInfoOut, error) {
	var profilePersonalInfo model.UserPersonalInfoOut
	err := r.db.QueryRow(ctx, "SELECT id,company, first_name, second_name from customer_personal_info WHERE id=$1", id).Scan(&profilePersonalInfo.Id, &profilePersonalInfo.Company, &profilePersonalInfo.FirstName, &profilePersonalInfo.SecondName)
	if err != nil {
		fmt.Println("user doesnt found")
		return profilePersonalInfo, nil
	}
	return profilePersonalInfo, err
}

func (r *userRepo) FetchProfileDeliveryInfo(ctx context.Context, id string) (model.UserDeliveryInfoOut, error) {
	var profileDeliveryInfo model.UserDeliveryInfoOut
	err := r.db.QueryRow(ctx, "SELECT id,phone_number, city, address from customer_delivery_info WHERE id = $1", id).Scan(&profileDeliveryInfo.Id, &profileDeliveryInfo.PhoneNumber, &profileDeliveryInfo.City, &profileDeliveryInfo.Address)
	if err != nil {
		fmt.Println("user doesnt found for delivery info")
		return profileDeliveryInfo, nil
	}
	return profileDeliveryInfo, nil
}

func (r *userRepo) RecordPersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	fmt.Println("Record new info for user", input.Id)
	fmt.Println(input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_personal_info (id,company, first_name, second_name) VALUES ($1, $2, $3, $4) RETURNING id, company, first_name,second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Id, &userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while inserting personal info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UpdatePersonalInfo(ctx context.Context, input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	fmt.Println("Update personal info for user")
	fmt.Println(input)
	err := r.db.QueryRow(ctx, "UPDATE customer_personal_info SET company =$2, first_name=$3, second_name=$4 WHERE id =$1 RETURNING company, first_name, second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		fmt.Println("error while updating delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) RecordDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	fmt.Println("Record new info for user", input.Id)
	fmt.Println(input)
	err := r.db.QueryRow(ctx, "INSERT INTO customer_delivery_info (id,phone_number, city, address) VALUES ($1, $2, $3, $4) RETURNING id, phone_number, city,address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.Id, &userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		fmt.Println("Error while inserting delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UpdateDeliveryInfo(ctx context.Context, input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	fmt.Println("Update delivery info for user")
	fmt.Println(input)
	err := r.db.QueryRow(ctx, "UPDATE customer_delivery_info SET phone_number=$2, city=$3, address=$4 WHERE id =$1 RETURNING phone_number, city, address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		fmt.Println("error while updating delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UserPasswordChange(ctx context.Context, input model.PasswordChange) (model.PasswordChange, error) {
	var check model.PasswordChange
	fmt.Println("чекаем пароль для юзера", input)
	err := r.db.QueryRow(ctx, "SELECT password FROM customer WHERE id =$1", input.Id).Scan(&check.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(check.OldPassword), []byte(input.OldPassword))
	if err != nil {
		fmt.Println("user writes wrong old password")
		return check, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 12)
	err = r.db.QueryRow(ctx, "UPDATE customer SET password = $1 WHERE id = $2 RETURNING password,id", hashedPassword, input.Id).Scan(&check.NewPassword, &check.Id)
	if err != nil {
		fmt.Println("ошибка")
		return check, err
	}
	return check, nil
}

func (r *userRepo) UserChangeEmail(ctx context.Context, input model.UserOrdinaryInfo) (model.UserOrdinaryInfo, error) {
	var user model.UserOrdinaryInfo
	var loginExists bool
	err := r.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM customer WHERE login =$1)", input.Login).Scan(&loginExists)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return user, err
	}
	if !loginExists {
		fmt.Println("login does not exist")
		return user, err
	}
	err = r.db.QueryRow(ctx, "UPDATE customer SET email = $1 WHERE login = $2 RETURNING login,email", input.Email, input.Login).Scan(&user.Login, &user.Email)
	if err != nil {
		fmt.Println("ошибка")
		return user, err
	}
	return user, nil
}

func (r *userRepo) GetAllUsers(ctx context.Context) ([]model.UserOrdinaryInfo, error) {
	fmt.Println("Getting all users")
	var users []model.UserOrdinaryInfo
	rows, err := r.db.Query(ctx, "SELECT id,login,password,email, registration_date,user_role FROM customer")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.UserOrdinaryInfo
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegistrationDate, &user.UserRole)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
