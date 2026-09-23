package users_transport_http

import (
	"net/http"
	"sync"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) GetUserProfile(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userIdJWT := ctx.Value("user_id")

	idStr, ok := userIdJWT.(string)
	if !ok {
		//NEED TO FIX//
		http.Error(rw, "invalid JWT token", http.StatusBadRequest)
		//NEED TO FIX//
		return
	}
	var user model.UserSummary

	profileCh := make(chan model.UserMerge)

	wg := new(sync.WaitGroup)

	wg.Add(3)

	go func() {
		defer wg.Done()
		data, err := h.usersService.FetchProfileInfo(r.Context(), idStr)

		profileCh <- model.UserMerge{
			Kind:  "profile_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.usersService.FetchProfilePersonalInfo(r.Context(), idStr)

		profileCh <- model.UserMerge{
			Kind:  "personal_info",
			Data:  data,
			Error: err,
		}
	}()

	go func() {
		defer wg.Done()
		data, err := h.usersService.FetchProfileDeliveryInfo(r.Context(), idStr)

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
			responseHandler.ErrorResponse(
				result.Error,
				"error when trying to get info from profile channel(fan-in)")
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

	responseHandler.ResponseWithJSON(http.StatusOK, user)
}
