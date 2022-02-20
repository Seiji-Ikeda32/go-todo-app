package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/db"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/handlers"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/repositories"
)

const Port = ":8080"

var tr = repositories.NewTodoRepository()
var tc = handlers.NewTodoHandler(tr)
var ro = handlers.NewRouter(tc)

func main() {
	db := db.OpenDB()
	defer db.Close()

	fmt.Println("Hello world")

	healthHandler := handlers.NewHealthHandler()

	http.HandleFunc("/health", healthHandler.ServeHTTP)

	http.HandleFunc("/todos/", ro.HandleTodoRequest)

	log.Fatal(http.ListenAndServe(Port, nil))
}
