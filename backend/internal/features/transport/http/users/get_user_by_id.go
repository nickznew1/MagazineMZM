package users_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetUserById(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	userIdJWT := ctx.Value("user_id")

	id, ok := userIdJWT.(string)
	if !ok {
		// NEED TO FIX //
		http.Error(rw, "invalid JWT token", http.StatusBadRequest)
		// NEED TO FIX //
		return
	}
	response, err := h.usersService.FetchProfileInfo(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when trying to fetch profile info")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
