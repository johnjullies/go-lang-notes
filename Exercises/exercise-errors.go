package main

import (
  "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}
  
func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return x, ErrNegativeSqrt(x)
  }
  z := 1.0
  temp := fmt.Sprintf("%.4f", z)
  for range make([]int, 10) {
    z -= (z*z - x) / (2*z)
    if check(z, temp) {
      return z, nil
    }
    temp = fmt.Sprintf("%.4f", z)
    // fmt.Printf("%d\t%.4f\n", i, z)
  }
  return z, nil
}

func check(z float64, temp string) bool {
  return fmt.Sprintf("%.4f", z) == temp
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
