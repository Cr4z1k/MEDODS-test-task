package service

import (
	"github.com/Cr4z1k/MEDODS-test-task/internal/core"
	"github.com/Cr4z1k/MEDODS-test-task/internal/repository"
	"github.com/Cr4z1k/MEDODS-test-task/pkg/auth"
)

type Tokens interface {
	GetTokens(guid string) (core.Tokens, error)
	RefreshTokens(refTokenBase64 string) (core.Tokens, error)
}

type Service struct {
	Tokens
}

func NewService(r *repository.Repository, t auth.TokenManager) *Service {
	return &Service{Tokens: NewTokenService(r.Tokens, t)}
}
