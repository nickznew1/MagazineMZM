package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) InsertPersonalInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("first insert for userInfo")
	var input model.UserPersonalInfo

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Wrong info for user")
		return
	}
	fmt.Println(input)
	newInfo, err := h.useCase.RecordPersonalInfo(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "error when record new info for user")
		return
	}
	RespondWithJSON(w, http.StatusCreated, newInfo)
}
