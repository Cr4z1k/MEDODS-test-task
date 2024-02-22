package service

import (
	"github.com/Cr4z1k/MEDODS-test-task/internal/repository"
	"github.com/Cr4z1k/MEDODS-test-task/pkg/auth"
)

type Service struct {
}

func NewService(r *repository.Repository, t *auth.Manager) *Service {
	return &Service{}
}
