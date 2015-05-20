package main


import (
  "fmt"
  "net/http"
  "github.com/jbuberel/example_nut/vendor/_nuts/github.com/gorilla/mux"
)

func HomeHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, world")
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.ListenAndServe(":8080", r)

}
