package repository

import (
	"database/sql"
	"fmt"
	"sync"

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

func (r *userRepo) GetUserById(input model.User) (model.User, error) {
	var user model.User
	fmt.Println("Getting user by ID")
	err := r.db.QueryRow("SELECT id,login, password FROM customer WHERE login =$1", input.Login).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		fmt.Println("user not found")
		return user, err
	}
	fmt.Println("get user with login ", user.Login)
	return user, nil
}

func (r *userRepo) CreateUser(input model.User) (model.User, error) {

	var customer model.User
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		fmt.Println("cant hash password")
		return customer, err
	}
	input.Password = string(hashedPassword)

	err = r.db.QueryRow(
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

func (r *userRepo) UserAuth(input model.User) (model.User, error) {
	fmt.Println("user authentication")
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

func (r *userRepo) FetchProfile(id string) (model.UserSummary, error) {
	var user model.UserSummary

	profileCh := make(chan model.UserMerge)

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		defer wg.Done()
		data, err := r.FetchProfileInfo(id)

		profileCh <- data{
			kind: "profile_info",
			data: data,
			err:  err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := r.FetchProfilePersonalInfo(id)

		profileCh <- data{
			kind: "personal_info",
			data: data,
			err:  err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := r.FetchProfileDeliveryInfo(id)

		profileCh <- data{
			kind: "delivery_info",
			data: data,
			err:  err,
		}
	}

	go func() {
		wg.Wait()
		close(profileCh)
	}()

	for result := range profileCh {
		if result.err != nil {
			return nil, err
		}

		switch result.kind {
		case "profile_info":
			if data, ok := result.data.(model.UserOrdinaryInfo); ok {
				user.UserOrdinary = &data
			}

		case "personal_info":
			if data, ok := result.data.(model.UserPersonalInfo); ok {
				user.UserPersonal = &data
			}

		case "delivery_info":
			if data, ok := result.data.(model.UserDeliveryInfo); ok {
				user.UserDelivery = &data
			}
		}
	}

	return user, nil
}

func (r *userRepo) FetchProfileInfo(id string) (model.User, error) {
	err := r.db.QueryRow("SELECT id,login,email,user_role from customer WHERE id =$1", id).Scan(&user.Id, &user.Login, &user.Email, &user.UserRole)
	if err != nil {
		fmt.Println("user doesnt found")
		return user, err
	}
	return user, nil
}

func (r *userRepo) FetchProfilePersonalInfo(id string) (model.UserPersonalInfo, error) {
	err := r.db.QueryRow("SELECT id,company, first_name, second_name from customer_personal_info WHERE id=$1", id).Scan(&userInfo.Id, &userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		fmt.Println("user doesnt found")
		return userInfo, nil
	}
	return userInfo, err
}

func (r *userRepo) FetchProfileDeliveryInfo(id string) (model.UserDeliveryInfo, error) {
	err := r.db.QueryRow("SELECT id,phone_number, city, address from customer_delivery_info WHERE id = $1", id).Scan(&userInfo.Id, &userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		fmt.Println("user doesnt found for delivery info")
		return userInfo, nil
	}
	return userInfo, nil
}

func (r *userRepo) RecordPersonalInfo(input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	fmt.Println("Record new info for user", input.Id)
	fmt.Println(input)
	err := r.db.QueryRow("INSERT INTO customer_personal_info (id,company, first_name, second_name) VALUES ($1, $2, $3, $4) RETURNING id, company, first_name,second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Id, &userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while inserting personal info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UpdatePersonalInfo(input model.UserPersonalInfo) (model.UserPersonalInfo, error) {
	var userInfo model.UserPersonalInfo
	fmt.Println("Update personal info for user")
	fmt.Println(input)
	err := r.db.QueryRow("UPDATE customer_personal_info SET company =$2, first_name=$3, second_name=$4 WHERE id =$1 RETURNING company, first_name, second_name",
		input.Id, input.Company, input.FirstName, input.SecondName).
		Scan(&userInfo.Company, &userInfo.FirstName, &userInfo.SecondName)
	if err != nil {
		fmt.Println("error while updating delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) RecordDeliveryInfo(input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	fmt.Println("Record new info for user", input.Id)
	fmt.Println(input)
	err := r.db.QueryRow("INSERT INTO customer_delivery_info (id,phone_number, city, address) VALUES ($1, $2, $3, $4) RETURNING id, phone_number, city,address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.Id, &userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		fmt.Println("Error while inserting delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UpdateDeliveryInfo(input model.UserDeliveryInfo) (model.UserDeliveryInfo, error) {
	var userInfo model.UserDeliveryInfo
	fmt.Println("Update delivery info for user")
	fmt.Println(input)
	err := r.db.QueryRow("UPDATE customer_delivery_info SET phone_number=$2, city=$3, address=$4 WHERE id =$1 RETURNING phone_number, city, address",
		input.Id, input.PhoneNumber, input.City, input.Address).
		Scan(&userInfo.PhoneNumber, &userInfo.City, &userInfo.Address)
	if err != nil {
		fmt.Println("error while updating delivery info")
		return userInfo, err
	}
	return userInfo, nil
}

func (r *userRepo) UserPasswordChange(input model.PasswordChange) (model.PasswordChange, error) {
	var check model.PasswordChange
	fmt.Println("чекаем пароль для юзера", input)
	err := r.db.QueryRow("SELECT password FROM customer WHERE id =$1", input.Id).Scan(&check.OldPassword)
	err = bcrypt.CompareHashAndPassword([]byte(check.OldPassword), []byte(input.OldPassword))
	if err != nil {
		fmt.Println("user writes wrong old password")
		return check, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), 12)
	err = r.db.QueryRow("UPDATE customer SET password = $1 WHERE id = $2 RETURNING password,id", hashedPassword, input.Id).Scan(&check.NewPassword, &check.Id)
	if err != nil {
		fmt.Println("ошибка")
		return check, err
	}
	return check, nil
}

func (r *userRepo) UserChangeEmail(input model.User) (model.User, error) {
	var user model.User
	var loginExists bool
	err := r.db.QueryRow("SELECT EXISTS (SELECT 1 FROM customer WHERE login =$1)", input.Login).Scan(&loginExists)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return user, err
	}
	if !loginExists {
		fmt.Println("login does not exist")
		return user, err
	}
	err = r.db.QueryRow("UPDATE customer SET email = $1 WHERE login = $2 RETURNING login,email", input.Email, input.Login).Scan(&user.Login, &user.Email)
	if err != nil {
		fmt.Println("ошибка")
		return user, err
	}
	return user, nil
}

func (r *userRepo) GetAllUsers() ([]model.User, error) {
	fmt.Println("Getting all users")
	var users []model.User
	rows, err := r.db.Query("SELECT id,login,password,email, registration_date,user_role FROM customer")
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.RegistrationDate, &user.UserRole)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
