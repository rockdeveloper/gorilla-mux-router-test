package main

import (
  "fmt",
  "github.com/gorilla/mux"
)


func Handler() {
    
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", Handler)
    http.Handle("/", r)
}