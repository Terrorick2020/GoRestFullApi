package service

import (
	"github.com/Terrorick2020/GoRestFullApi/internal"

	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list internal.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
