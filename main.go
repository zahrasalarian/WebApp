package main

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type","text/html")
  if r.URL.Path == "/"{
    fmt.Fprint(w, "<h1>Hello</h1>")
  }else if r.URL.Path == "/tbbt"{
    fmt.Fprint(w, "Check it out <a href=\"https://en.wikipedia.org/wiki/The_Big_Bang_Theory\">The big bang theory</a>.")
  }else{
    fmt.Fprint(w, "<h1>Sorry, you're lost</h1>")
  }
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main()  {
  router := httprouter.New()
  router.GET("/hello/:name", Hello)

  // http.HandleFunc("/", handlerFunc)
  http.ListenAndServe(":2000", router)
}
