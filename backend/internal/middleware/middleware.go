package middleware

import (
	"context"
	"encoding/json"
	"myWebApp/backend/pkg/auth"
	"net/http"
	"strings"
)

type Manager struct {
	tokenManager auth.TokenManager
}

func NewManager(tokenManager auth.TokenManager) *Manager {
	return &Manager{
		tokenManager: tokenManager,
	}
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	respondWithJSON(w, statusCode, map[string]string{"error": message})
}

func (manager *Manager) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "didnt get token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {

			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		userId, err := manager.tokenManager.Parse(tokenString)

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
