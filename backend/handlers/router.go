package handlers

import (
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
	"github.com/labstack/echo"
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
	e.GET("/todos", th.GetTodos)
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
