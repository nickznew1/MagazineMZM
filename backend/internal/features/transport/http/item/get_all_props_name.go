package item_transport_http

import (
	"net/http"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (h *ItemHTTPHandler) GetAllPropsName(
	rw http.ResponseWriter,
	r *http.Request) {

	ctx := r.Context()

	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	response, err := h.itemService.GetAllPropsName(ctx)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get all props")
		return
	}

	responseHandler.ResponseWithJSON(http.StatusOK, response)
}
