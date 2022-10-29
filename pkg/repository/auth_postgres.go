package repository

import (
	"fmt"
	"github.com/ArtemFed/todo-app-test"
	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthService {
	return &AuthService{db: db}
}

func (r *AuthService) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
