package handlers

import (
	"net/http"
	"strings"
)

type Router interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	th TodoHandler
}

// 引数にTodoControllerを受け取り、Router構造体のポインタを返却
func NewRouter(th TodoHandler) Router {
	return &router{th}
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	userId := strings.TrimPrefix(r.URL.Path, "/todos/")
	switch {
	case r.Method == "GET" && userId == "":
		ro.th.GetTodos(w, r)
	case r.Method == "GET" && userId != "":
		ro.th.GetTodo(w, r)
	case r.Method == "POST":
		ro.th.PostTodo(w, r)
	case r.Method == "PUT":
		ro.th.PutTodo(w, r)
	case r.Method == "DELETE":
		ro.th.DeleteTodo(w, r)
	}
}
