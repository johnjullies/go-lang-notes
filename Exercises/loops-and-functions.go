package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  temp := z
  for i := range make([]int, 10) {
    z -= (z*z - x) / (2*z)
    fmt.Printf("%d\t%.4f\n", i, z)
    check := fmt.Sprintf("%.4f", z) == fmt.Sprintf("%.4f", temp)
    if check {
      return z
    }
    temp = z
  }
  return z
}

func main() {
  fmt.Println(Sqrt(2))
}
