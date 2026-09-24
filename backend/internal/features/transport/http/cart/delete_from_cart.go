package cart_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *CartHTTPHandler) DeleteFromCart(rw http.ResponseWriter, r *http.Request) {
	var input model.Cart

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode request body")
		return
	}
	response, err := h.cartService.DeleteFromCart(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to complete response")
		return
	}
	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
