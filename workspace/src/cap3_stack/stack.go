package main

import (
  "errors"
  "fmt"
)

func main() {
  stack := Stack{}
  fmt.Println("Stack created with size", stack.Size())
  fmt.Println("Empty?", stack.Empty())

  stack.Add("Go")
  stack.Add(12345)
  stack.Add(3.14)
  stack.Add("End")
  fmt.Println("Size after stack 4 values: ", stack.Size())
  fmt.Println("Empty", stack.Empty())

  for !stack.Empty(){
    v, _ := stack.Remove()
    fmt.Println("Removing..", v)
    fmt.Println"Size:", stack.Size())
    fmt.Println("Empty?", stack.Empty())
  }

  _, err := stack.Remove()
  if err != nil {
    fmt.Println("Error: ", err)
  }
}
