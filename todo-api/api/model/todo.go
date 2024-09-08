package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string
	Title       string
	Email       string
	Description string
	DueDate     time.Time
	Done        bool
}

func NewTodo(title, description, email string, dueDate time.Time) (*Todo, error) {
	if len(title) == 0 || len(description) == 0 || len(email) == 0 {
		return nil, errors.New("Invalid fields")
	}
	todo := &Todo{
		Title:       title,
		Email:       email,
		Description: description,
		DueDate:     dueDate,
		Done:        false,
	}
	todo.ID = uuid.NewString()
	return todo, nil
}

type TodoRepository interface {
	Save(*Todo) error
	FindById(string) (*Todo, error)
	List() ([]Todo, error)
	Finish(id string) error
}

type TodoRepositoryMemoryImpl struct {
	todos []Todo
}

func NewTodoRepositoryMemoryImpl() *TodoRepositoryMemoryImpl {
	todos := make([]Todo, 0, 0)
	return &TodoRepositoryMemoryImpl{todos}
}

func (r *TodoRepositoryMemoryImpl) Save(todo *Todo) error {
	r.todos = append(r.todos, *todo)
	return nil
}

func (r *TodoRepositoryMemoryImpl) List() ([]Todo, error) {
	return r.todos, nil
}

func (r *TodoRepositoryMemoryImpl) FindById(id string) (*Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Not found")
}

func (r *TodoRepositoryMemoryImpl) Finish(id string) error {
	for idx, todo := range r.todos {
		if todo.ID == id {
			todo.Done = true
			r.todos[idx] = todo
			return nil
		}
	}

	return errors.New("Not found")
}
