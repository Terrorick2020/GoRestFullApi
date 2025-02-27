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
	Create(userId, listId int, item internal.TodoItem) (int, error)
	GetAll(userId, listId int) ([]internal.TodoItem, error)
	GetById(userId, itemId int) (internal.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input internal.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
