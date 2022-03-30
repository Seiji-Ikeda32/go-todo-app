package handlers

import (
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// 引数にTodoHandlerを受け取り、Router構造体のポインタを返却
func NewRouter() *echo.Echo {
	tr := repositories.NewTodoRepository()
	th := NewTodoHandler(tr)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/todos", th.GetTodos)
	e.GET("/todos/:id", th.GetTodo)
	e.POST("/todos", th.PostTodo)
	e.PUT("/todos/:id", th.PutTodo)
	e.DELETE("/todos/:id", th.DeleteTodo)

	return e
}
