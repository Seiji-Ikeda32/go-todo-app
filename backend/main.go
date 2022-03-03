package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/handlers"
)

const Port = ":8080"

// var tr = repositories.NewTodoRepository()
// var th = handlers.NewTodoHandler(tr)

// var ro = handlers.NewRouter(tc)

func main() {

	fmt.Println("Hello world")

	healthHandler := handlers.NewHealthHandler()

	http.HandleFunc("/health", healthHandler.ServeHTTP)

	// http.HandleFunc("/todos/", ro.HandleTodoRequest)
	router := handlers.NewRouter()
	router.Logger.Fatal(router.Start(Port))

	log.Fatal(http.ListenAndServe(Port, nil))
}
