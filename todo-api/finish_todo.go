package main

type FinishTodo struct {
	repository TodoRepository
}

func NewFinishTodo(repository TodoRepository) *FinishTodo {
	return &FinishTodo{repository}
}

func (ft *FinishTodo) Execute(todoId string) error {
	return ft.repository.Finish(todoId)
}
