package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
	"github.com/labstack/echo"
)

type TodoHandler interface {
	GetTodo(c echo.Context) error
	GetTodos(c echo.Context) error
	PostTodo(c echo.Context) error
	PutTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandler struct {
	tr repositories.TodoRepository
}

// 引数にTodoRepositoryを受け取り、TodoHandler構造体のポインタを返却
func NewTodoHandler(tr repositories.TodoRepository) TodoHandler {
	return &todoHandler{tr}
}

func (th *todoHandler) GetTodo(c echo.Context) (err error) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	todo, err := th.tr.GetTodo(todoId)
	if err != nil {
		log.Println(err)
		return
	}

	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandler) GetTodos(c echo.Context) (err error) {
	todos, err := th.tr.GetTodos()
	if err != nil {
		log.Println(err)
		return
	}

	return c.JSON(http.StatusOK, todos)
}

func (th *todoHandler) PostTodo(c echo.Context) (err error) {
	c.Request().Header.Set("Content-Type", echo.MIMEApplicationJSONCharsetUTF8)

	var todoRequest models.TodoRequest
	if err = c.Bind(&todoRequest); err != nil {
		log.Println(err)
		return
	}

	todo := models.Todo{
		Title:       todoRequest.Title,
		Discription: todoRequest.Discription,
		DueTime:     todoRequest.DueTime,
		CreatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	}

	err = th.tr.CreateTodo(todo)
	if err != nil {
		log.Println(err)
		return
	}

	return c.String(http.StatusOK, "created successfully")
}

func (th *todoHandler) PutTodo(c echo.Context) (err error) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	var todoRequest models.TodoRequest
	if err = c.Bind(&todoRequest); err != nil {
		log.Println(err)
		return
	}

	todo := models.Todo{
		Id:          todoId,
		Title:       todoRequest.Title,
		Discription: todoRequest.Discription,
		IsCompleted: todoRequest.IsCompleted,
		DueTime:     todoRequest.DueTime,
		UpdatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	}

	err = th.tr.UpdateTodo(todo)
	if err != nil {
		log.Println(err)
		return
	}

	return c.String(http.StatusOK, "updated successfully")
}

func (th *todoHandler) DeleteTodo(c echo.Context) (err error) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	err = th.tr.DeleteTodo(todoId)
	if err != nil {
		log.Println(err)
		return
	}

	return c.String(http.StatusOK, "deleted successfully")
}
