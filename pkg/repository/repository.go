package repository

import (
	"github.com/ArtemFed/todo-app-test"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type UserList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

//
//type TodoItem interface {
//}

type Repository struct {
	Authorization
	UserList
	//TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserList:      NewUsersListPostgres(db),
	}
}
