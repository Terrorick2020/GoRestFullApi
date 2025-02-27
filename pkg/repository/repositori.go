package repository

import (
	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user internal.User) (int, error)
	GetUser(username, password string) (internal.User, error)
}

type TodoList interface {
	Create(userId int, list internal.TodoList) (int, error)
	GetAll(userId int) ([]internal.TodoList, error)
	GetById(userId, listId int) (internal.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input internal.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item internal.TodoItem) (int, error)
	GetAll(userId, listId int) ([]internal.TodoItem, error)
	GetById(userId, itemId int) (internal.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input internal.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
