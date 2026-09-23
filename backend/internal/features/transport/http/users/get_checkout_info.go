package users_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetCheckoutInfo(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userIdJWT := ctx.Value("user_id")

	idStr, ok := userIdJWT.(string)
	if !ok {
		// NEED TO FIX //
		http.Error(rw, "invalid JWT token", http.StatusBadRequest)
		// NEED TO FIX //
		return
	}

	responsePersonal, err := h.usersService.FetchProfilePersonalInfo(r.Context(), idStr)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when trying to get profile personal info")
	}
	responseDelivery, err := h.usersService.FetchProfileDeliveryInfo(r.Context(), idStr)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when trying to get profile delivery info")
	}
	responseProfile, err := h.usersService.FetchProfileInfo(r.Context(), idStr)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when trying to get full profile info")
	}

	responseHandler.ResponseWithJSON(http.StatusCreated, map[string]interface{}{
		"user":     responseProfile,
		"personal": responsePersonal,
		"delivery": responseDelivery,
	})
}
