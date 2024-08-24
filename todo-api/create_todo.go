package main

import (
	"net/http"
	"time"
)

type CreateTodoDto struct {
	Title       string    `json:"title"`
	Descritpion string    `json:"description"`
	Email       string    `json:"email"`
	DueDate     time.Time `json:"due_date"`
}

type CreateTodo struct {
	repository TodoRepository
}

func NewCreateTodo(repository TodoRepository) *CreateTodo {
	return &CreateTodo{repository}
}

func (ct *CreateTodo) Execute(dto *CreateTodoDto) error {
	todo, err := NewTodo(dto.Title, dto.Descritpion, dto.Email, dto.DueDate)
	if err != nil {
		return NewApiError(err.Error(), http.StatusBadRequest)
	}

	if err := ct.repository.Save(todo); err != nil {
		return NewApiError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
