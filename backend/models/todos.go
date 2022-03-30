package models

import (
	"database/sql"
)

type Todo struct {
	Id          int          `json:"id"`
	Title       string       `json:"title"`
	Discription string       `json:"discription"`
	IsCompleted bool         `json:"isCompleted"`
	DueTime     sql.NullTime `json:"dueTime"`
	CreatedAt   sql.NullTime `json:"createdAt"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
}

type TodoRequest struct {
	Title       string       `json:"title"`
	Discription string       `json:"discription"`
	IsCompleted bool         `json:"isCompleted"`
	DueTime     sql.NullTime `json:"dueTime"`
	CreatedAt   sql.NullTime `json:"createdAt"`
	UpdatedAt   sql.NullTime `json:"updatedAt"`
}

type TodosResponse struct {
	Todos []Todo `json:"todos"`
}
