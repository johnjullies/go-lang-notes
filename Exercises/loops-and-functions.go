package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  z := 1.0
  temp := fmt.Sprintf("%.4f", z)
  for i := range make([]int, 10) {
    z -= (z*z - x) / (2*z)
    if check(z, temp) {
      return z
    }
    temp = fmt.Sprintf("%.4f", z)
    fmt.Printf("%d\t%.4f\n", i, z)
  }
  return z
}

func check(z float64, temp string) bool {
  return fmt.Sprintf("%.4f", z) == temp
}

func main() {
  fmt.Println(Sqrt(2))
}
