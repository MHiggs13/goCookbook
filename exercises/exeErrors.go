package main

import (
  "fmt"
)

type ErrNegativeSqrt float64


func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("Cannot find square root of negative (%v)", float64(e))
}

func Sqrt(x float64) (float64, error) {
  // if x is neg, check sqrt on absolute value
  if x<0 {
    return x, ErrNegativeSqrt(x)
  }

  // start checking at 0
  var z float64 = 0
  for ; z*z <= x; z++ {
    if z*z == float64(x) {
      return z, nil;
    }
  }
  // prints approx. if z*z is not exactly x
  fmt.Print("Approx:")
  return z, nil;

}

func main() {
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}
