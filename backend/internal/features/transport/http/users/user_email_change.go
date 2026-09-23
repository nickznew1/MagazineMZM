package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UserEmailChange(
	rw http.ResponseWriter,
	r *http.Request) {

	var input model.UserOrdinaryInfo

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when decode request body")
		return
	}
	response, err := h.usersService.UserChangeEmail(r.Context(), input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when change user email")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusCreated, response)
}
