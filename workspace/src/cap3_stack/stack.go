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
    fmt.Println("Size:", stack.Size())
    fmt.Println("Empty?", stack.Empty())
  }

  _, err := stack.Remove()
  if err != nil {
    fmt.Println("Error: ", err)
  }
}

type Stack struct {
  values []interface{}
}

func (stack Stack)Size()int {
  return len(stack.values)
}

func (stack Stack)Empty() bool{
  return stack.Size() == 0
}

func (stack *Stack) Add(value interface{}){
  stack.values = append(stack.values, value)
}

func (stack *Stack)Remove()(interface{}, error){
  if stack.Empty(){
    return nil, errors.New("Empty stack!")
  }
  value := stack.values[stack.Size()-1]
  stack.values = stack.values[:stack.Size() - 1]
  return value, nil
}
