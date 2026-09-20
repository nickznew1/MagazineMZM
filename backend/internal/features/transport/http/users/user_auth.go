package users_transport_http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *UsersHTTPHandler) UserAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user authentification")
	var input model.UserOrdinaryInfo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Неверно отправленные данные")
		return
	}
	userAuth, err := h.useCase.UserAuth(r.Context(), input)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	tokenId := strconv.Itoa(userAuth.Id)
	user, err := h.auth.NewJWT(tokenId)
	/*user, err := h.store.GetSessionByLogin(input.Login)*/
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "oshibka polycheniya id pri logine")
		return
	}
	RespondWithJSON(w, http.StatusOK, map[string]string{
		"access_token": user,
	})
}
