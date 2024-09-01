package main

import (
	"encoding/json"
	"log"
)

type FinishTodo struct {
	repository TodoRepository
	producer   Producer
}

type KafkaMessageDto struct {
	Email  string `json:"email"`
	TodoId string `json:"todo_id"`
}

func NewFinishTodo(repository TodoRepository, producer Producer) *FinishTodo {
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

	if err := ft.producer.SendMessage(b); err != nil {
		log.Printf("Error sending message to kafka %v\n", err)
		return
	}
}
