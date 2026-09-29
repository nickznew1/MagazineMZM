package item_transport_http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ItemHTTPHandler) GetSpecById(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	idQuery := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(idQuery)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"parsing query to integer")
		return
	}
	specId, err := h.itemService.GetSpecById(r.Context(), idInt)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"get specification id")
		return
	}

	responseHandler.ResponseWithJSON(http.StatusOK, specId)
}
