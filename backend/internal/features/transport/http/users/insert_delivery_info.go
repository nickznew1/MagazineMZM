package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) InsertDeliveryInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for DeliveryInfo")
	var input model.UserDeliveryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "delivery info wrong")
		return
	}
	userInfo, err := h.useCase.RecordDeliveryInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when record delivery info")
		return
	}
	RespondWithJSON(w, http.StatusCreated, userInfo)
}
