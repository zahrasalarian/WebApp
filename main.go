package main

import (
  "fmt"
  "net/http"
  // "github.com/julienschmidt/httprouter"
  "github.com/gorilla/mux"
)


func home(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Hello</h1>")
}

func tbbt(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "Check it out <a href=\"https://en.wikipedia.org/wiki/The_Big_Bang_Theory\">The big bang theory</a>.")
}
func faq(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>This is my FAQ page</h1>")
}

func main()  {
  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/tbbt", tbbt)
  r.HandleFunc("/faq", faq)
  http.ListenAndServe(":2000", r)
}
