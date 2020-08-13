package main

import (
  "fmt"
  "net/http"
  "html/template"
  // "github.com/julienschmidt/httprouter"
  "github.com/gorilla/mux"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  if err := homeTemplate.Execute(w, nil); err != nil{

  }
}

func tbbt(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "Check it out <a href=\"https://en.wikipedia.org/wiki/The_Big_Bang_Theory\">The big bang theory</a>.")
}

func main()  {
  var err error
  homeTemplate, err = template.ParseFiles("views/home.gohtml")
  if err != nil{
    panic(err)
  }
  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/tbbt", tbbt)
  http.ListenAndServe(":2000", r)
}
