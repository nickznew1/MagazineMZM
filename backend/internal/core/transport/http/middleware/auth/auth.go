package core_middleware_auth

import (
	"context"
	"net/http"
	"strings"

	core_logger "github.com/nickznew1/MagazineMZM/backend/internal/core/logger"
	core_http_response "github.com/nickznew1/MagazineMZM/backend/internal/core/transport/http/response"
)

func (m *Manager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		
		logger := core_logger.FromContext(ctx)

		responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

		authHeader := r.Header.Get("Authorization")

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userId, err := m.tokenManager.Parse(tokenString)

		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"error when trying parse JWT token")
			return
		}
		ctx = context.WithValue(r.Context(), "user_id", userId)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})

}
