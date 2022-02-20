package handlers

import "net/http"

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
		ro.th.GetTodos(w, r)
	case "POST":
		ro.th.PostTodo(w, r)
	}
}
