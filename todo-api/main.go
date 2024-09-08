package main

import (
	"log"

	"github.com/danilomarques1/todo-api/api"
	"github.com/danilomarques1/todo-api/api/model"
	"github.com/danilomarques1/todo-api/api/producer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	todoRepository := model.NewTodoRepositoryMemoryImpl()
	queueProducer, err := producer.NewProducer(producer.QUEUE_PRODUCER)
	if err != nil {
		log.Fatal(err)
	}
	defer queueProducer.Close()
	kafkaProducer, err := producer.NewProducer(producer.KAFKA_PRODUCER)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaProducer.Close()

	api.NewTodoApi(e, kafkaProducer, queueProducer, todoRepository).Register()

	e.Logger.Fatal(e.Start(":5000"))
}
