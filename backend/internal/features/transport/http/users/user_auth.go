package users_transport_http

import (
	"encoding/json"
	"net/http"
	"strconv"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UserAuth(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	logger := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	var input model.UserOrdinaryInfo

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when decode request body")
		return
	}
	userAuth, err := h.usersService.UserAuth(r.Context(), input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when auth user")
		return
	}
	tokenId := strconv.Itoa(userAuth.Id)

	user, err := h.auth.NewJWT(tokenId)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"error when creating new JWT token")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, map[string]string{
		"access_token": user,
	})
}
