package main

import (
  "fmt"
  "html"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", TodoShow)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request){
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  //fmt.Fprintf(w, "Todo show: ", todoId)
  fmt.Fprintf(w, "Todo show: %q", todoId)
  //fmt.Fprintf(w, "Todo show: %q", html.EscapeString(todoId))
}
