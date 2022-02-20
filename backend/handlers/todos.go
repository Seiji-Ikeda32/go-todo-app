package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
)

type TodoHandler interface {
	PostTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	tr repositories.TodoRepository
}

func NewTodoHandler(tr repositories.TodoRepository) TodoHandler {
	return &todoHandler{tr}
}

func (th *todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := th.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []models.Todo
	for _, v := range todos {
		todoResponses = append(todoResponses, models.Todo{
			Id:          v.Id,
			Title:       v.Title,
			Discription: v.Discription,
			IsCompleted: v.IsCompleted,
			DueTime:     v.DueTime,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}

	var todosResponse models.TodosResponse
	todosResponse.Todos = todoResponses

	res, _ := json.Marshal(todosResponse.Todos)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (th *todoHandler) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest models.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := models.Todo{Title: todoRequest.Title, Discription: todoRequest.Discription}

	id, err := th.tr.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}
