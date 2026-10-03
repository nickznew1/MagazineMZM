package applications_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
)

func (h *ApplicationHTTPHandler) CreateApplication(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var input model.Application

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode request body")
		return
	}

	response, err := h.applicationService.CreateApplication(ctx, input)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create application")
		return
	}

	responseHandler.ResponseWithJSON(
		http.StatusCreated,
		map[string]interface{}{
			"application_id": response,
		},
	)

}
