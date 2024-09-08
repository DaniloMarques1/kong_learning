package service

import (
	"net/http"

	"github.com/danilomarques1/todo-api/api/model"
)

type ListTodo struct {
	repository model.TodoRepository
}

func NewListTodo(repository model.TodoRepository) *ListTodo {
	return &ListTodo{repository}
}

func (lt *ListTodo) Execute() ([]model.Todo, error) {
	todos, err := lt.repository.List()
	if err != nil {
		return nil, model.NewApiError(err.Error(), http.StatusInternalServerError)
	}
	return todos, nil
}
