package main

import "fmt"

func main(){

  a, b := 0, 10
  for a < b{
    a += 1
    fmt.Println(a)
  }

 //using range
 numbers := []int{1, 2, 3, 4, 5}
 for i := range numbers{
   numbers[i] *= 2
 }
 fmt.Println(numbers)
}
