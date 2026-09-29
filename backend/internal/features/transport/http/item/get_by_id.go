package item_transport_http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ItemHTTPHandler) GetById(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	idString := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(idString)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"parsing url query")
		return
	}
	id, err := h.itemService.GetById(ctx, idInt)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"get item id")
		return
	}

	responseHandler.ResponseWithJSON(http.StatusOK, id)

}
