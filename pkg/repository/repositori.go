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
