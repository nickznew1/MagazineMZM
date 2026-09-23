package core_http_response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HTTPResponseHandler struct {
	log *slog.Logger

	rw http.ResponseWriter
}

func NewHTTPResponseHandler(log *slog.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) ResponseWithJSON(statusCode int, responseBody any) {
	h.rw.Header().Set("Content-Type", "application/json")
	h.rw.WriteHeader(statusCode)
	if err := json.NewEncoder(h.rw).Encode(responseBody); err != nil {
		//LOGGER
	}
}

func (h *HTTPResponseHandler) ResponseWithError(statusCode int, err error, message string) {
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": message,
		"error":   err.Error(),
	}

	h.ResponseWithJSON(statusCode,
		response,
	)

}
