package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
)

type TodoHandler interface {
	GetTodo(w http.ResponseWriter, r *http.Request)
	GetTodos(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoHandler struct {
	tr repositories.TodoRepository
}

func NewTodoHandler(tr repositories.TodoRepository) TodoHandler {
	return &todoHandler{tr}
}

func (th *todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	todo, err := th.tr.GetTodo(todoId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Println(todo)

	res, _ := json.Marshal(todo)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (th *todoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := th.tr.GetTodos()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []models.Todo
	for _, todo := range todos {
		todoResponses = append(todoResponses, models.Todo{
			Id:          todo.Id,
			Title:       todo.Title,
			Discription: todo.Discription,
			IsCompleted: todo.IsCompleted,
			DueTime:     todo.DueTime,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
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

	todo := models.Todo{
		Title:       todoRequest.Title,
		Discription: todoRequest.Discription,
		DueTime:     todoRequest.DueTime,
		CreatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	}

	id, err := th.tr.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (th *todoHandler) PutTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var todoRequest models.TodoRequest
	json.Unmarshal(body, &todoRequest)

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
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (th *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = th.tr.DeleteTodo(todoId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}
