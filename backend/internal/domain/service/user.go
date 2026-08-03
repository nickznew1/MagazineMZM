package service

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"github.com/nickznew1/MagazineMZM/backend/pkg/auth"
	"net/http"
	"strconv"
)

type UserService struct {
	useCase     usecase.UserUseCase
	auth        auth.TokenManager
	cartUseCase *usecase.CartUseCase
}

func NewUserService(c *usecase.UserUseCase, a auth.TokenManager) *UserService {
	return &UserService{
		useCase: *c,
		auth:    a,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func (h *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	user, err := h.useCase.CreateUser(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "неверные данные")
		return
	}
	tokenId := strconv.Itoa(user.Id)
	token, err := h.auth.NewJWT(tokenId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "oshibka pri sozdanii sessii")
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{
		"access_token": token,
	})
}

func (h *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user")
	userIdJWT := r.Context().Value("user_id")
	idStr, ok := userIdJWT.(string)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "invalid user id format")
		return
	}
	userId, err := h.useCase.FetchProfileInfo(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	respondWithJSON(w, http.StatusOK, userId)
}

func (h *UserService) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUserProfile")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "invalid user id format")
		return
	}
	var user model.UserSummary

	profileCh := make(chan model.UserMerge)

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfileInfo(idStr)

		profileCh <- model.UserMerge{
			Kind:  "profile_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfilePersonalInfo(idStr)

		profileCh <- model.UserMerge{
			Kind:  "personal_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfileDeliveryInfo(idStr)

		profileCh <- model.UserMerge{
			Kind:  "delivery_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		wg.Wait()
		close(profileCh)
	}()

	for result := range profileCh {
		if result.Error != nil {
			respondWithError(w, http.StatusBadRequest, "invalid data")
		}

		switch result.Kind {
		case "profile_info":
			if data, ok := result.Data.(model.UserOrdinaryInfoOut); ok {
				user.UserOrdinary = &data
			}

		case "personal_info":
			if data, ok := result.Data.(model.UserPersonalInfoOut); ok {
				user.UserPersonal = &data
			}

		case "delivery_info":
			if data, ok := result.Data.(model.UserDeliveryInfoOut); ok {
				user.UserDelivery = &data
			}
		}
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (h *UserService) InsertPersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for userInfo")
	var input model.UserPersonalInfo

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Wrong info for user")
		return
	}
	fmt.Println(input)
	newInfo, err := h.useCase.RecordPersonalInfo(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error when record new info for user")
		return
	}
	respondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) UpdatePersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user personal")
	var input model.UserPersonalInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong data for update user personal")
		return
	}
	newInfo, err := h.useCase.UpdatePersonalInfo(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error when try to update personal info")
		return
	}
	respondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) InsertDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for DeliveryInfo")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "delivery info wrong")
		return
	}
	userInfo, err := h.useCase.RecordDeliveryInfo(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error when record delivery info")
		return
	}
	respondWithJSON(w, http.StatusCreated, userInfo)
}

func (h *UserService) UpdateDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user delivery")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "wrong data for update user delivery")
		return
	}
	newInfo, err := h.useCase.UpdateDeliveryInfo(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "error when try to update delivery info")
		return
	}
	respondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) UserAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user authentification")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверно отправленные данные")
		return
	}
	userAuth, err := h.useCase.UserAuth(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	tokenId := strconv.Itoa(userAuth.Id)
	user, err := h.auth.NewJWT(tokenId)
	/*user, err := h.store.GetSessionByLogin(input.Login)*/
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "oshibka polycheniya id pri logine")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{
		"access_token": user,
	})
}

func (h *UserService) UserPasswordChange(w http.ResponseWriter, r *http.Request) {
	var changes model.PasswordChange
	if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
		fmt.Println(changes)
		respondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newPassword, err := h.useCase.UserPasswordChange(changes)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusCreated, newPassword)
}

func (h *UserService) UserEmailChange(w http.ResponseWriter, r *http.Request) {
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newEmail, err := h.useCase.UserChangeEmail(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusCreated, newEmail)
}

func (h *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	Users, err := h.useCase.GetAllUsers()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusCreated, Users)
}

func (h *UserService) GetCheckoutInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get checkout ")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		respondWithError(w, http.StatusBadRequest, "invalid user id format")
		return
	}
	userInfo, _ := h.useCase.FetchProfilePersonalInfo(idStr)
	userDelivery, _ := h.useCase.FetchProfileDeliveryInfo(idStr)
	///*cart, _ := h.cartUseCase.GetCart(idStr)*/
	user, _ := h.useCase.FetchProfileInfo(idStr)

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"user":     user,
		"personal": userInfo,
		"delivery": userDelivery,
	})
}
