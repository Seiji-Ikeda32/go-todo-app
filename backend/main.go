package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/db"
	"github.com/Seiji-Ikeda32/go-todo-app/backend/handlers"
)

const Port = ":8080"

func mainHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello World")
	if err != nil {
		return
	}
}

func main() {
	db := db.OpenDB()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println("faild connect database", err)
		return
	} else {
		fmt.Println("success connect database")
	}
	fmt.Println("Hello world")

	http.HandleFunc("/", mainHandler)

	healthHandler := handlers.NewHealthHandler()

	http.HandleFunc("/health", healthHandler.ServeHTTP)

	log.Fatal(http.ListenAndServe(Port, nil))
}
