package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type CreateTodoDto struct {
	Title       string    `json:"title"`
	Descritpion string    `json:"description"`
	Email       string    `json:"email"`
	DueDate     time.Time `json:"due_date"`
}

type SchedulerMessageDto struct {
	Email            string    `json:"email"`
	NotificationDate time.Time `json:"notifcationDate"`
}

type CreateTodo struct {
	repository TodoRepository
	producer   Producer
}

func NewCreateTodo(repository TodoRepository, producer Producer) *CreateTodo {
	return &CreateTodo{repository, producer}
}

func (ct *CreateTodo) Execute(dto *CreateTodoDto) error {
	todo, err := NewTodo(dto.Title, dto.Descritpion, dto.Email, dto.DueDate)
	if err != nil {
		return NewApiError(err.Error(), http.StatusBadRequest)
	}

	if err := ct.repository.Save(todo); err != nil {
		return NewApiError(err.Error(), http.StatusInternalServerError)
	}

	schedulerMessage := SchedulerMessageDto{Email: todo.Email, NotificationDate: todo.DueDate}
	go ct.sendMessageToQueue(schedulerMessage)
	return nil
}

func (ct *CreateTodo) sendMessageToQueue(schedulerMessage SchedulerMessageDto) {
	b, err := json.Marshal(schedulerMessage)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	if err := ct.producer.SendMessage(b); err != nil {
		log.Printf("There was an error pushing a message to the queue %v\n", err)
		return
	}
}
