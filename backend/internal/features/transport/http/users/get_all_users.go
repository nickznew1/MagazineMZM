package users_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) GetAllUsers(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	response, err := h.usersService.GetAllUsers(ctx)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"service error")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
