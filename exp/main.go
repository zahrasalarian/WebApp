package main

import(
  "os"
  "html/template"
)

type User struct{
  Name string
  Dog string
}

func main()  {
  t, err := template.ParseFiles("hello.gohtml")
  if err != nil{
    panic(err)
  }

  data := User{
    Name: "Amy Farrah Fowler",
    Dog: "Cinnamon",
  }

  err = t.Execute(os.Stdout,data)
  if err != nil{
    panic(err)
  }
}
