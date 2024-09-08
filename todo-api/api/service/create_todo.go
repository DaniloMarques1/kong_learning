package service

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/danilomarques1/todo-api/api/model"
	"github.com/danilomarques1/todo-api/api/producer"
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
	repository model.TodoRepository
	qProducer  producer.Producer
}

func NewCreateTodo(repository model.TodoRepository, qProducer producer.Producer) *CreateTodo {
	return &CreateTodo{repository, qProducer}
}

func (ct *CreateTodo) Execute(dto *CreateTodoDto) error {
	todo, err := model.NewTodo(dto.Title, dto.Descritpion, dto.Email, dto.DueDate)
	if err != nil {
		return model.NewApiError(err.Error(), http.StatusBadRequest)
	}

	if err := ct.repository.Save(todo); err != nil {
		return model.NewApiError(err.Error(), http.StatusInternalServerError)
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

	if err := ct.qProducer.SendMessage(b); err != nil {
		log.Printf("There was an error pushing a message to the queue %v\n", err)
		return
	}
}
