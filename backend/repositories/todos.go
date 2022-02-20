package repositories

import (
	"log"
	"time"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/db"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
)

type TodoRepository interface {
	GetTodos() (todos []models.Todo, err error)
	CreateTodo(todo models.Todo) (id int, err error)
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodos() (todos []models.Todo, err error) {
	db := db.OpenDB()
	todos = []models.Todo{}

	rows, err := db.Query("SELECT id, title, discription, isCompleted, due_time, created_at, updated_at FROM todos ORDER BY id DESC")
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		todo := models.Todo{}
		err = rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.Discription,
			&todo.IsCompleted,
			&todo.DueTime,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			log.Println(err)
			return
		}
		todos = append(todos, todo)
	}

	return
}

func (tr *todoRepository) CreateTodo(todo models.Todo) (id int, err error) {
	db := db.OpenDB()
	cmd := `insert into todos (
		title,
		discription,
		due_time,
		created_at) values (?, ?, ?, ?)`
	_, err = db.Exec(cmd,
		todo.Title,
		todo.Discription,
		todo.DueTime,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	err = db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)
	return
}
