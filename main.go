package main

import (  
 "fmt"
 "reflect"
//  uuid "github.com/satori/go.uuid"
) 

func main() {
  // fmt.Println("Hello World")
  t := "test"
  fmt.Printf("%+v", reflect.TypeOf(t))
  
}

// type MyJoe struct {
//   id string
// }

// func NewMyJoe() MyJoe {
  
// }