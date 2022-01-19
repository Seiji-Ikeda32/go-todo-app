package main

import (
    "fmt"
    "log"
    "net/http"

    . "../internal/handlers/api"
)

const Port = ":8080"

func mainHandler(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprint(w, "Hello World")
    if err != nil {
        return
    }
}

func main() {
    fmt.Println("Hello world")

    http.HandleFunc("/", mainHandler)

    healthHandler := NewHealthHandler()

    http.HandleFunc("/health", healthHandler.ServeHTTP)

    log.Fatal(http.ListenAndServe(Port, nil))
}
