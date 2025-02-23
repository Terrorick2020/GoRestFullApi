package repository

import (
	"github.com/Terrorick2020/GoRestFullApi/internal"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user internal.User) (int, error) {
	return 0, nil
}