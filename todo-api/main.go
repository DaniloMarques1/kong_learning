package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiResponseErrorDto struct {
	ErrorMessage string `json:"error_message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	todoRepository := NewTodoRepositoryMemoryImpl()
	e.POST("/todo", func(c echo.Context) error {
		// amqp://fitz:fitz@localhost:5672
		producer, err := NewProducer("amqp://fitz:fitz@rabbitmq:5672")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ApiResponseErrorDto{ErrorMessage: err.Error()})
		}
		createTodo := NewCreateTodo(todoRepository, producer)
		createTodoDto := &CreateTodoDto{}
		if err := c.Bind(createTodoDto); err != nil {
			return c.JSON(http.StatusBadRequest, ApiResponseErrorDto{ErrorMessage: "Invalid body"})
		}
		if err := createTodo.Execute(createTodoDto); err != nil {
			apiError := err.(*ApiError)
			return c.JSON(apiError.Code, ApiResponseErrorDto{ErrorMessage: apiError.Msg})
		}

		return c.NoContent(http.StatusNoContent)
	})

	e.GET("/todo", func(c echo.Context) error {
		listTodo := NewListTodo(todoRepository)
		todos, err := listTodo.Execute()
		if err != nil {
			apiError := err.(*ApiError)
			return c.JSON(apiError.Code, ApiResponseErrorDto{ErrorMessage: apiError.Msg})
		}
		return c.JSON(http.StatusOK, todos)
	})

	e.GET("/todo", func(c echo.Context) error {
		listTodo := NewListTodo(todoRepository)
		todos, err := listTodo.Execute()
		if err != nil {
			apiError := err.(*ApiError)
			return c.JSON(apiError.Code, ApiResponseErrorDto{ErrorMessage: apiError.Msg})
		}
		return c.JSON(http.StatusOK, todos)
	})

	e.PUT("/todo/:todo_id", func(c echo.Context) error {
		todoId := c.Param("todo_id")
		finishTodo := NewFinishTodo(todoRepository)
		if err := finishTodo.Execute(todoId); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":5000"))
}
