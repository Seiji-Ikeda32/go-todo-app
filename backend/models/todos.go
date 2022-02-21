package models

import (
	"database/sql"
)

type Todo struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Discription string       `json:"discription"`
	IsCompleted bool         `json:"is_completed"`
	DueTime     sql.NullTime `json:"due_time"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type TodoRequest struct {
	Title       string       `json:"title"`
	Discription string       `json:"discription"`
	IsCompleted bool         `json:"is_completed"`
	DueTime     sql.NullTime `json:"due_time"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type TodosResponse struct {
	Todos []Todo `json:"todos"`
}
