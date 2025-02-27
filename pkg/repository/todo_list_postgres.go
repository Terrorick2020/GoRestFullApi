package repository

import (
	"strings"
	"fmt"

	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (t *TodoListPostgres) Create(userId int, list internal.TodoList) (int, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, descripion) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
        tx.Rollback()
        return 0, err
    }

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
        tx.Rollback()
        return 0, err
    }

	return id, tx.Commit()
}

func (t *TodoListPostgres) GetAll(userId int) ([]internal.TodoList, error) {
	var list []internal.TodoList

	query := fmt.Sprintf(
		"SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.user_id WHERE ul.user_id = $1",
		todoListsTable,
		userListsTable,
	)
	err := t.db.Select(&list, query, userId)

	return list, err
}

func (t *TodoListPostgres) GetById(userId, listId int) (internal.TodoList, error) {
	var list internal.TodoList

	query := fmt.Sprintf(
		`SELECT tl.id, tl.title, tl.description FROM %s tl
		INNER JOIN %s ul ON tl.id = ul.user_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, userListsTable,
	)
	err := t.db.Get(&list, query, userId, listId)

	return list, err
}

func (t *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s tl USING %s ul ON tl.id = ul.user_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, userListsTable,
	)
	_, err := t.db.Exec(query, userId, listId)


	return err
}

func (t *TodoListPostgres) Update(userId, listId int, input internal.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id = $%d",
		todoListsTable, setQuery, userListsTable, argId, argId + 1,
	)

	args = append(args, listId, userId)

	_, err := t.db.Exec(query, args...)

	return err
}
