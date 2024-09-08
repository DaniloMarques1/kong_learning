package api

import (
	"net/http"

	"github.com/danilomarques1/todo-api/api/model"
	"github.com/danilomarques1/todo-api/api/producer"
	"github.com/danilomarques1/todo-api/api/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoApi struct {
	g              *echo.Group
	kProducer      producer.Producer
	qProducer      producer.Producer
	todoRepository model.TodoRepository
}

func NewTodoApi(e *echo.Echo, kProducer producer.Producer, qProducer producer.Producer, todoRepository model.TodoRepository) *TodoApi {
	g := e.Group("/todo")
	return &TodoApi{g, kProducer, qProducer, todoRepository}
}

func (t *TodoApi) Register() {
	t.g.POST("/todo", func(c echo.Context) error {
		createTodo := service.NewCreateTodo(t.todoRepository, t.qProducer)
		createTodoDto := &service.CreateTodoDto{}
		if err := c.Bind(createTodoDto); err != nil {
			return model.ResponseError(c, model.NewApiError("Invalid body", http.StatusBadRequest))
		}
		if err := createTodo.Execute(createTodoDto); err != nil {
			return model.ResponseError(c, err)
		}

		return c.NoContent(http.StatusNoContent)
	})

	t.g.GET("/todo", func(c echo.Context) error {
		listTodo := service.NewListTodo(t.todoRepository)
		todos, err := listTodo.Execute()
		if err != nil {
			return model.ResponseError(c, err)
		}
		return c.JSON(http.StatusOK, todos)
	})

	t.g.GET("/todo", func(c echo.Context) error {
		listTodo := service.NewListTodo(t.todoRepository)
		todos, err := listTodo.Execute()
		if err != nil {
			return model.ResponseError(c, err)
		}
		return c.JSON(http.StatusOK, todos)
	})

	t.g.PUT("/todo/finish/:todo_id", func(c echo.Context) error {
		todoId := c.Param("todo_id")
		if _, err := uuid.Parse(todoId); err != nil {
			return model.ResponseError(c, model.NewApiError("Invalid todo id", http.StatusBadRequest))
		}
		finishTodo := service.NewFinishTodo(t.todoRepository, t.kProducer)
		if err := finishTodo.Execute(todoId); err != nil {
			return model.ResponseError(c, err)
		}

		return c.NoContent(http.StatusNoContent)
	})
}
