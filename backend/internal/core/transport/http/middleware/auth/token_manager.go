package core_middleware_auth

import (
	"crypto/rand"
	"encoding/base64"
	"log/slog"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "ggm9WQnw5wpBLYR7MZ0cuMA24BRAFodaQpmUkJuM6QO"
)

type TokenManager interface {
	NewJWT(userId string) (string, error)
	Parse(accessToken string) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey   string
	tokenManager TokenManager
	log          *slog.Logger
}

func NewManager(log *slog.Logger) (*Manager, error) {
	return &Manager{
		signingKey: signingKey,
		log:        log,
	}, nil
}

func (m *Manager) NewJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Subject:   userId,
	})
	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			m.log.Error("unexpected signing method: ", token.Header["alg"])
			return nil, err
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		m.log.Error("error get user claims from token", err)
		return "", err
	}

	return claims["sub"].(string), nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		m.log.Error("error get user refresh token", err)
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
