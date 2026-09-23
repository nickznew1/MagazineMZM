package core_http_response

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	core_errors "github.com/nickznew1/MagazineMZM/backend/internal/core/errors"
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

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		loggerFunc func(string, ...any)
	)
	
	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		loggerFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
		loggerFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
		loggerFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		loggerFunc = h.log.Error
	}
	loggerFunc(msg, slog.String("error: ", error.Error(err)))

	h.ResponseWithError(statusCode, err, msg)
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
