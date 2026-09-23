package users_transport_http

import (
	"encoding/json"
	"net/http"
	"strconv"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) CreateUser(
	rw http.ResponseWriter,
	r *http.Request) {

	var input model.UserOrdinaryInfo

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when decode request")
		return
	}
	response, err := h.usersService.CreateUser(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"service error")
	}

	tokenId := strconv.Itoa(response.Id)

	token, err := h.auth.NewJWT(tokenId)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when creating JWT token")
	}
	responseHandler.ResponseWithJSON(http.StatusCreated, map[string]string{
		"access_token": token,
	})
}
