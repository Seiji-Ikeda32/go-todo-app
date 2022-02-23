package handlers

import (
	"net/http"
	"path"
)

type Router interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	th TodoHandler
}

func NewRouter(th TodoHandler) Router {
	return &router{th}
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if path.Base(r.URL.Path) == "/" {
			ro.th.GetTodos(w, r)
		} else {
			ro.th.GetTodo(w, r)
		}
	case "POST":
		ro.th.PostTodo(w, r)
	case "PUT":
		ro.th.PutTodo(w, r)
	case "DELETE":
		ro.th.DeleteTodo(w, r)
	}
}
