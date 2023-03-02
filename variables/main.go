package main

import "fmt"

func main() {
  var a = "initial"
  fmt.Println(a) // string

  var b, c = 1, 2
  fmt.Println(b, c) // if all variables have only one type, you can use this declaration format

  var d = true
  fmt.Println(d) // bool

  var e int 
  fmt.Println(e) // use this format without value, return 0 int

  f := "apple"
  fmt.Println(f) // you can use this shortcut to declare variables
}
