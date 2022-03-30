package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/handlers"
)

const Port = ":8080"

func main() {

	fmt.Println("Hello world")

	healthHandler := handlers.NewHealthHandler()

	http.HandleFunc("/health", healthHandler.ServeHTTP)

	router := handlers.NewRouter()
	router.Logger.Fatal(router.Start(Port))

	log.Fatal(http.ListenAndServe(Port, nil))
}
