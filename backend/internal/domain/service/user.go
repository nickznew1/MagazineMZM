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

func (h *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	user, err := h.useCase.CreateUser(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "неверные данные")
		return
	}
	tokenId := strconv.Itoa(user.Id)
	token, err := h.auth.NewJWT(tokenId)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "oshibka pri sozdanii sessii")
	}
	RespondWithJSON(w, http.StatusCreated, map[string]string{
		"access_token": token,
	})
}

func (h *UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user")
	userIdJWT := r.Context().Value("user_id")
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "invalid user id format")
		return
	}
	userId, err := h.useCase.FetchProfileInfo(r.Context(), idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusOK, userId)
}

func (h *UserService) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getUserProfile")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "invalid user id format")
		return
	}
	var user model.UserSummary

	profileCh := make(chan model.UserMerge)

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfileInfo(r.Context(), idStr)

		profileCh <- model.UserMerge{
			Kind:  "profile_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfilePersonalInfo(r.Context(), idStr)

		profileCh <- model.UserMerge{
			Kind:  "personal_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.useCase.FetchProfileDeliveryInfo(r.Context(), idStr)

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
			RespondWithError(w, http.StatusBadRequest, "invalid data")
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

	RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserService) InsertPersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for userInfo")
	var input model.UserPersonalInfo

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Wrong info for user")
		return
	}
	fmt.Println(input)
	newInfo, err := h.useCase.RecordPersonalInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when record new info for user")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) UpdatePersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user personal")
	var input model.UserPersonalInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "wrong data for update user personal")
		return
	}
	newInfo, err := h.useCase.UpdatePersonalInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when try to update personal info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) InsertDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for DeliveryInfo")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "delivery info wrong")
		return
	}
	userInfo, err := h.useCase.RecordDeliveryInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when record delivery info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, userInfo)
}

func (h *UserService) UpdateDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user delivery")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "wrong data for update user delivery")
		return
	}
	newInfo, err := h.useCase.UpdateDeliveryInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when try to update delivery info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}

func (h *UserService) UserAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user authentification")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверно отправленные данные")
		return
	}
	userAuth, err := h.useCase.UserAuth(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	tokenId := strconv.Itoa(userAuth.Id)
	user, err := h.auth.NewJWT(tokenId)
	/*user, err := h.store.GetSessionByLogin(input.Login)*/
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "oshibka polycheniya id pri logine")
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"access_token": user,
	})
}

func (h *UserService) UserPasswordChange(w http.ResponseWriter, r *http.Request) {
	var changes model.PasswordChange
	if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
		fmt.Println(changes)
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newPassword, err := h.useCase.UserPasswordChange(r.Context(), changes)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newPassword)
}

func (h *UserService) UserEmailChange(w http.ResponseWriter, r *http.Request) {
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newEmail, err := h.useCase.UserChangeEmail(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newEmail)
}

func (h *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	Users, err := h.useCase.GetAllUsers()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, Users)
}

func (h *UserService) GetCheckoutInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get checkout ")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "invalid user id format")
		return
	}
	userInfo, _ := h.useCase.FetchProfilePersonalInfo(r.Context(), idStr)
	userDelivery, _ := h.useCase.FetchProfileDeliveryInfo(r.Context(), idStr)
	user, _ := h.useCase.FetchProfileInfo(r.Context(), idStr)

	RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"user":     user,
		"personal": userInfo,
		"delivery": userDelivery,
	})
}
