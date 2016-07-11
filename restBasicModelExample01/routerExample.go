package main

import (
  "log"
  "net/http"
  "time"
   "encoding/json"

  "github.com/gorilla/mux"
)

type Todo struct {
  Name string
  Completed bool
  Due time.Time
}

type Todos []Todo


func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/todos", TodoIndex)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  todos := Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetuo"},
  }

  json.NewEncoder(w).Encode(todos)

}
