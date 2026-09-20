package users_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UserEmailChange(w http.ResponseWriter, r *http.Request) {
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newEmail, err := h.useCase.UserChangeEmail(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newEmail)
}
