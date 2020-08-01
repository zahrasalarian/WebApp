package main

import (
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprint(w, "<h1>fg</h1>")
}

func main()  {
  http.HandleFunc("/", handlerFunc)
  http.ListenAndServe(":2000", nil)
}
