package main

import(
  "fmt"
  "os"
  "strconv"
)

func main(){

  input := os.Args[1:]
  numbers := make([]int, len(input))

  for i, n := range input {
    number, err := strconv.Atoi(n)
    if err != nil {
      fmt.Printf("%s not a valid number! \n")
      os.Exit(1)
    }
    numbers[i] = number
  }

  fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int)[]int {
  if len(numbers)<= 1{
    return numbers
  }

  n := make([]int, len(numbers))
  copy(n, numbers)

  pivotIndex := len(n) / 2
  pivot := n[pivotIndex]
  n = append(n[:pivotIndex],n[pivotIndex+1:]...)

  lowers, greaters := partition(n, pivot)

  return append(append(quicksort(lowers), pivot), quicksort(greaters)...)
}

func partition(numbers []int, pivot int)(lowers []int, greaters []int) {
  for _, n := range numbers{
    if n <= pivot {
      lowers = append(lowers, n)
    }else{
      greaters = append(greaters, n)
    }
  }
  return lowers, greaters
}
