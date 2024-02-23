package service

import (
	"encoding/base64"
	"time"

	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"github.com/Cr4z1k/MEDODS-test-task/internal/repository"
	"github.com/Cr4z1k/MEDODS-test-task/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

const (
	ttl = time.Hour // time.Minute * 15
)

type TokensService struct {
	r repository.Tokens
	t auth.TokenManager
}

func NewTokenService(r repository.Tokens, t auth.TokenManager) *TokensService {
	return &TokensService{
		r: r,
		t: t,
	}
}

func (s *TokensService) GetTokens(guid string) (core.Tokens, error) {
	refToken, err := s.t.NewRefreshToken()
	if err != nil {
		return core.Tokens{}, err
	}

	accToken, err := s.t.NewAccessToken(guid, ttl)
	if err != nil {
		return core.Tokens{}, err
	}

	hashedRefToken, err := bcrypt.GenerateFromPassword([]byte(refToken), bcrypt.DefaultCost)
	if err != nil {
		return core.Tokens{}, err
	}

	err = s.r.GetTokens(guid, string(hashedRefToken))
	if err != nil {
		return core.Tokens{}, err
	}

	refTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refToken))

	tokens := core.Tokens{
		AccToken: accToken,
		RefToken: refTokenBase64,
	}

	return tokens, nil
}
