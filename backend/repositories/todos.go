package repositories

import (
	"log"
	"time"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/db"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) (id int, err error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) CreateTodo(todo models.Todo) (id int, err error) {
	db := db.OpenDB()
	cmd := `insert into todos (
		title,
		discription,
		created_at) values (?, ?, ?)`
	_, err = db.Exec(cmd,
		todo.Title,
		todo.Discription,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	err = db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}
