package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UpdatePersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user personal")
	var input model.UserPersonalInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "wrong data for update user personal")
		return
	}
	newInfo, err := h.useCase.UpdatePersonalInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when try to update personal info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}
