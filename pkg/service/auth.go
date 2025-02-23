package service

import (
	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user internal.User) (int, error) {
	return s.repo.CreateUser(user)
}
