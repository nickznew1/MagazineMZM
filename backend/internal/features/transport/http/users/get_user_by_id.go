package users_transport_http

import (
	"fmt"
	"net/http"
)

func (h *UsersHTTPHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get user")
	userIdJWT := r.Context().Value("user_id")
	idStr, ok := userIdJWT.(string)
	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "invalid user id format")
		return
	}
	userId, err := h.useCase.FetchProfileInfo(r.Context(), idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusOK, userId)
}
