package users_transport_http

import (
	"encoding/json"
	"fmt"
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
		responseHandler.ResponseWithError()
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
