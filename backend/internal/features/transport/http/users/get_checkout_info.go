package users_transport_http

import (
	"fmt"
	"net/http"
)

func (h *UsersHTTPHandler) GetCheckoutInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get checkout ")
	userIdJWT := r.Context().Value("user_id")
	fmt.Println("userIdJWT ", userIdJWT)
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "invalid user id format")
		return
	}
	userInfo, _ := h.useCase.FetchProfilePersonalInfo(r.Context(), idStr)
	userDelivery, _ := h.useCase.FetchProfileDeliveryInfo(r.Context(), idStr)
	user, _ := h.useCase.FetchProfileInfo(r.Context(), idStr)

	RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"user":     user,
		"personal": userInfo,
		"delivery": userDelivery,
	})
}
