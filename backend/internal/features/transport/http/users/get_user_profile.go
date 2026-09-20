package users_transport_http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
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
