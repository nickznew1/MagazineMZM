package cart_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *CartHTTPHandler) GetCart(rw http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userIdJWT := r.Context().Value("user_id")

	idStr, ok := userIdJWT.(string)

	if !ok {
		// NEED TO FIX //
		http.Error(rw, "invalid user id format", http.StatusForbidden)
		// NEED TO FIX //
		return
	}
	response, err := h.cartService.GetCart(ctx, idStr)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to complete response")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
