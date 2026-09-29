package item_transport_http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *ItemHTTPHandler) ChangeVisible(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var input model.Item

	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode response body")
		return
	}
	response, err := h.itemService.ChangeVisible(ctx, input.Visible, id)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to change visible for item")
		return
	}

	responseHandler.ResponseWithJSON(http.StatusOK, response)

}
