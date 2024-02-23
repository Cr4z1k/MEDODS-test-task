package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	NewAccessToken(guid string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	signingKey string
}

func NewManager(signingKey string) (*Manager, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewAccessToken(guid string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		Subject:   guid,
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *Manager) NewRefreshToken() (string, error) {
	a := make([]byte, 24)

	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)

	if _, err := rnd.Read(a); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", a), nil
}
