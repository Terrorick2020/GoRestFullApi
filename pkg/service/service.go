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
	GetAll(userId int) ([]internal.TodoList, error)
	GetById(userId, listId int) (internal.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input internal.UpdateListInput) error
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
