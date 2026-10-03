package applications_transport_http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ApplicationHTTPHandler) GetApplication(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	idStr := chi.URLParam(r, "id")

	userIdJWT := r.Context().Value("user_id")

	response, err := h.applicationService.GetApplication(r.Context(), idStr, userIdJWT.(string))
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get application")
	}

	responseHandler.ResponseWithJSON(
		http.StatusOK,
		response,
	)
}
