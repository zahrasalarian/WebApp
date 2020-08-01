package main

import (
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type","text/html")
  if r.URL.Path == "/"{
    fmt.Fprint(w, "<h1>Hello</h1>")
  }else if r.URL.Path == "/tbbt"{
    fmt.Fprint(w, "Check it out <a href=\"https://en.wikipedia.org/wiki/The_Big_Bang_Theory\">The big bang theory</a>.")
  }
}

func main()  {
  http.HandleFunc("/", handlerFunc)
  http.ListenAndServe(":2000", nil)
}
