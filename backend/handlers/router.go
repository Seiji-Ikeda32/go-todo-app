package handlers

import (
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// type Router interface {
// 	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
// }

// type router struct {
// 	th TodoHandler
// }

// 引数にTodoHandlerを受け取り、Router構造体のポインタを返却
func NewRouter() *echo.Echo {
	tr := repositories.NewTodoRepository()
	th := NewTodoHandler(tr)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/todos", th.GetTodos)
	e.GET("/todos/:id", th.GetTodo)
	e.POST("/todos", th.PostTodo)
	e.PUT("/todos/:id", th.PutTodo)
	e.DELETE("/todos/:id", th.DeleteTodo)

	return e
}

// func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
// 	userId := strings.TrimPrefix(r.URL.Path, "/todos")
// 	switch {
// 	case r.Method == "GET" && userId != "":
// 		ro.th.GetTodo(w, r)
// 	case r.Method == "POST":
// 		ro.th.PostTodo(w, r)
// 	case r.Method == "PUT":
// 		ro.th.PutTodo(w, r)
// 	case r.Method == "DELETE":
// 		ro.th.DeleteTodo(w, r)
// 	}
// }

// func InitRouting(e *echo.Echo) {
// 	tr := repositories.NewTodoRepository()
// 	th := NewTodoHandler(tr)

// 	e.GET("/todos", th.GetTodos)
// }
