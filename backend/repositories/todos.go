package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/db"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
)

type TodoRepository interface {
	GetTodo(id int) (todo models.Todo, err error)
	GetTodos() (todos []models.Todo, err error)
	CreateTodo(todo models.Todo) (err error)
	UpdateTodo(todo models.Todo) (err error)
	DeleteTodo(id int) (err error)
}

type todoRepository struct{}

// TodoRepository構造体のポインタを返却
func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodo(id int) (todo models.Todo, err error) {
	db := db.OpenDB()
	defer db.Close()

	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	todo = models.Todo{}
	// pointerで渡した型に取得したデータをマッピング
	err = row.Scan(
		&todo.Id,
		&todo.Title,
		&todo.Discription,
		&todo.IsCompleted,
		&todo.DueTime,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("指定したtodoは存在しません")
		} else {
			log.Println(err)
		}
	}

	return
}

func (tr *todoRepository) GetTodos() (todos []models.Todo, err error) {
	db := db.OpenDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todos ORDER BY id DESC")
	if err != nil {
		log.Println(err)
		return
	}

	todos = []models.Todo{}

	for rows.Next() {
		todo := models.Todo{}
		// pointerで渡した型に取得したデータをマッピング
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

func (tr *todoRepository) CreateTodo(todo models.Todo) (err error) {
	db := db.OpenDB()
	defer db.Close()

	cmd := `insert into todos (
		title,
		discription,
		due_time,
		created_at) values (?, ?, ?, ?)`

	_, err = db.Exec(cmd,
		todo.Title,
		todo.Discription,
		todo.DueTime,
		time.Now(),
	)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (tr *todoRepository) UpdateTodo(todo models.Todo) (err error) {
	db := db.OpenDB()
	cmd := `UPDATE todos SET
	    title = ?,
		discription = ?,
		IsCompleted = ?,
		due_time = ?,
		updated_at = ? WHERE id = ?`
	_, err = db.Exec(cmd,
		todo.Title,
		todo.Discription,
		todo.IsCompleted,
		todo.DueTime,
		time.Now(),
		todo.Id)
	return
}

func (tr *todoRepository) DeleteTodo(id int) (err error) {
	db := db.OpenDB()
	defer db.Close()
	_, err = db.Exec("DELETE FROM todos WHERE id = ?", id)
	return
}
