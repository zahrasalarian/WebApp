package main

import (
  // "fmt"
  "net/http"
  "html/template"
  // "github.com/julienschmidt/httprouter"
  "github.com/gorilla/mux"
)

var(
  homeTemplate *template.Template
  tbbtTemplate *template.Template
)
func home(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  if err := homeTemplate.Execute(w,nil); err != nil{
    panic(err)
  }
}

func tbbt(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "text/html")
  if err := tbbtTemplate.Execute(w,nil); err != nil{
    panic(err)
  }
}

func main()  {
  var err error
  homeTemplate, err = template.ParseFiles("views/home.gohtml")
  if err != nil{
    panic(err)
  }
  tbbtTemplate, err = template.ParseFiles("views/tbbt.gohtml")
  if err != nil{
    panic(err)
  }
  r := mux.NewRouter()
  r.HandleFunc("/", home)
  r.HandleFunc("/tbbt", tbbt)
  http.ListenAndServe(":2000", r)
}
