package service

import (
	"encoding/json"
	"log"

	"github.com/danilomarques1/todo-api/api/model"
	"github.com/danilomarques1/todo-api/api/producer"
)

type FinishTodo struct {
	repository model.TodoRepository
	p          producer.Producer
}

type KafkaMessageDto struct {
	Email  string `json:"email"`
	TodoId string `json:"todo_id"`
}

func NewFinishTodo(repository model.TodoRepository, producer producer.Producer) *FinishTodo {
	return &FinishTodo{repository, producer}
}

func (ft *FinishTodo) Execute(todoId string) error {
	todo, err := ft.repository.FindById(todoId)
	if err != nil {
		return err
	}

	if err := ft.repository.Finish(todoId); err != nil {
		return err
	}
	kafkaMessage := KafkaMessageDto{Email: todo.Email, TodoId: todo.ID}
	go ft.sendMessageToKafka(kafkaMessage)
	return nil
}

func (ft *FinishTodo) sendMessageToKafka(kafkaMessageDto KafkaMessageDto) {
	b, err := json.Marshal(kafkaMessageDto)
	if err != nil {
		log.Printf("Error parsing message %v\n", err)
		return
	}

	log.Printf("Sending message to kafka %v\n", string(b))

	if err := ft.p.SendMessage(b); err != nil {
		log.Printf("Error sending message to kafka %v\n", err)
		return
	}
}
