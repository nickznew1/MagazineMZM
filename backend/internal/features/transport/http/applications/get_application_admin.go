package applications_transport_http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ApplicationHTTPHandler) GetApplicationForAdmin(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	id := chi.URLParam(r, "id")

	response, err := h.applicationService.GetApplicationForAdmin(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get application for admin",
		)
	}
	responseHandler.ResponseWithJSON(
		http.StatusOK,
		response,
	)
}
