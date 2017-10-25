package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)


func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    fmt.Fprint(w, "<h1>gorilla mux simple test!</h1>");
}


func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", Handler)

    http.ListenAndServe(":3000", r)
}
