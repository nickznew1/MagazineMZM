package applications_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ApplicationHTTPHandler) GetAllApplicationsForUser(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userIdJWT := r.Context().Value("user_id")
	response, err := h.applicationService.GetAllApplicationsForUser(r.Context(), userIdJWT.(string))
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get all applications for user",
		)
		return
	}
	responseHandler.ResponseWithJSON(
		http.StatusOK,
		response,
	)
}
