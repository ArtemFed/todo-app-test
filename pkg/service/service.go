package service

import (
	"github.com/ArtemFed/todo-app-test"
	"github.com/ArtemFed/todo-app-test/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UsersList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

//type TodoItem interface {
//}

type Service struct {
	Authorization
	UsersList
	//TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UsersList:     NewUsersListService(repos.UserList),
	}
}
