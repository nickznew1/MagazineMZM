package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UserPasswordChange(w http.ResponseWriter, r *http.Request) {
	var changes model.PasswordChange
	if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
		fmt.Println(changes)
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newPassword, err := h.useCase.UserPasswordChange(r.Context(), changes)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newPassword)
}
