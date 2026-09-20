package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UpdateDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update user delivery")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "wrong data for update user delivery")
		return
	}
	newInfo, err := h.useCase.UpdateDeliveryInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when try to update delivery info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}
