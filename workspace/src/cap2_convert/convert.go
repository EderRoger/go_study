package main

import (
  "fmt"
  "os"
  "strconv"
)

func main(){
  if len(os.Args) < 3 {
    fmt.Println("Use: convert <values> <unity>")
    os.Exit(1)
  }

  originUnity := os.Args[len(os.Args)-1]
  originValues := os.Args[1 : len(os.Args)-1]

  var destinyUnity string

  if originUnity == "celsius" {
    destinyUnity = "fahrenheit"
  }else if originUnity == "quilometers" {
    destinyUnity = "miles"
  } else {
    fmt.Printf("%s not a know unity!", destinyUnity)
    os.Exit(1)
  }

  for i, v := range originValues {
    originValue, err := strconv.ParseFloat(v, 64)
    if err != nil {
      fmt.Printf("Value of %s in position %d not a valid number\n", v , i)
      os.Exit(1)
    }

    var destinyValue float64

    if originUnity == "celsius" {
      destinyValue = originValue*1.8+32
    }else{
      destinyValue = originValue / 1.60934
    }

    fmt.Printf("%.2f %s = %.2f %s\n", originValue, originUnity, destinyValue, destinyUnity)
  }
}
