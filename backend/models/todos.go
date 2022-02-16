package models

import "time"

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Discription string    `json:"discription"`
	IsCompleted bool      `json:"is_completed"`
	DueTime     time.Time `json:"due_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Discription string `json:"discription"`
}
