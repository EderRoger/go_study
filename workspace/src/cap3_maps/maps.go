package main

import (
  "fmt"
  "os"
  "strings"
)

func main(){
  words := os.Args[1:]
  statistics := getStatistics(words)
  show(statistics)
}

func getStatistics(words []string)map[string]int{
  statistics := make(map[string]int)

  for _, word := range words {
    initial := strings.ToUpper(string(word[0]))
    count, founded := statistics[initial]
    if founded {
      statistics[initial] = count + 1
    }else{
      statistics[initial] = 1
    }
  }
  return statistics
}

func show(statistics map[string]int){
  fmt.Println("Word counter begining no first word:")

  for initial, count := range statistics{
    fmt.Printf("%s = %d\n", initial, count)
  }
}
