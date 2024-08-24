package main

import "net/http"

type ListTodo struct {
	repository TodoRepository
}

func NewListTodo(repository TodoRepository) *ListTodo {
	return &ListTodo{repository}
}

func (lt *ListTodo) Execute() ([]Todo, error) {
	todos, err := lt.repository.List()
	if err != nil {
		return nil, NewApiError(err.Error(), http.StatusInternalServerError)
	}
	return todos, nil
}
