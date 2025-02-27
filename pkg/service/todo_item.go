package service

import (
	"github.com/Terrorick2020/GoRestFullApi/pkg/repository"
	"github.com/Terrorick2020/GoRestFullApi/internal"
)

type internalItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *internalItemService {
	return &internalItemService{repo: repo, listRepo: listRepo}
}

func (s *internalItemService) Create(userId, listId int, item internal.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *internalItemService) GetAll(userId, listId int) ([]internal.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *internalItemService) GetById(userId, itemId int) (internal.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *internalItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *internalItemService) Update(userId, itemId int, input internal.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
