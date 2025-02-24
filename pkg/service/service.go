package service

import (
	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
)

type Authorization interface {
	CreateUser(user internal.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list internal.TodoList) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
