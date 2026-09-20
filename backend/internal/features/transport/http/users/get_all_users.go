package users_transport_http

import "net/http"

func (h *UsersHTTPHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	Users, err := h.useCase.GetAllUsers(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	RespondWithJSON(w, http.StatusCreated, Users)
}
