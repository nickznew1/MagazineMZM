package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) InsertPersonalInfo(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	var input model.UserPersonalInfo

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when decode request body")
		return
	}

	// NEED TO RENAME METHOD //
	response, err := h.usersService.RecordPersonalInfo(r.Context(), input)
	// NEED TO RENAME METHOD //

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when insert personal info")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
