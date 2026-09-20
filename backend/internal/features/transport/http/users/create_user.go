package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
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
