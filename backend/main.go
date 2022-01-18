package main

import (
    "fmt"
    "log"
    "net/http"
)

const Port = ":8080"

type MainHandler struct{}

func (h *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
    fmt.Println("Hello world")

    log.Fatal(http.ListenAndServe(Port, &MainHandler{}))

}
