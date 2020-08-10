package main

import(
  "os"
  "html/template"
  // "fmt"
)

type User struct{
  Name string
  Dog DogInf
}
type DogInf struct{
    Name string
    Age int
}
type CastsInf struct {
      CastsName []string
 }

func main()  {

  // casts := []string{"Johnny Galecki","Jim Parsons","Kaley Cuoco","Simon Helberg","Kunal Nayyar","Mayim Bialik","Melissa Rauch"}
  t, err := template.ParseFiles("hello.gohtml")
  if err != nil{
    panic(err)
  }

  data := User{
    Name: "Amy Farrah Fowler",
    Dog: DogInf{
      Name: "Cinnamon",
      Age: 3,
    },
  }
  // castsInf := CastsInf{
  //   CastsName: casts,
  // }

  err = t.Execute(os.Stdout,data)
  if err != nil{
    panic(err)
  }
  // err = t.Execute(os.Stdout,castsInf)
  // if err != nil{
  //   panic(err)
  // }
  // fmt.Println(casts)
}
