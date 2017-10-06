package main

import (
  "fmt"
  "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  n := fmt.Sprint(float64(e))
  return fmt.Sprintf("cannot Sqrt negative number: %v", n)
}

func Sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, ErrNegativeSqrt(x)
  }
  var z float64 = 1.0
  for {
    nextZ := z - (z * z - x) / (2 * z)
    if math.Abs(nextZ - z) < 0.000000000000001 {
      break
    }
    z = nextZ
  }
  return z, nil
}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
