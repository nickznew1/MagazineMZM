package item_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *ItemHTTPHandler) GetItemId(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var input model.Item

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode response body")
		return
	}
	id, err := h.itemService.GetItemId(ctx, input)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get item id")
		return
	}

	responseHandler.ResponseWithJSON(http.StatusOK, id)
}
