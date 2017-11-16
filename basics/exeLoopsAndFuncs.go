package main

import (
  "fmt"
)

// As a way to play with functions and loops, let's impolement a square root function:
// given a number x, we want to find the number z for which z*z is most near x

func Sqrt(x float64) float64 {
  // if x is neg, check sqrt on absolute value
  if x<0 {
    Sqrt(-x)
  }

  // start checking at 0
  var z float64 = 0
  for ; z*z <= x; z++ {
    if z*z == float64(x) {
      return z;
    }
  }
  // prints approx. if z*z is not exactly x
  fmt.Print("Approx:")
  return z;

}

func main() {
  var a,b,c,d,e,f float64 = 2,4,61,125,-100,1
  fmt.Println("sqrt(",a,") : ",Sqrt(a))
  fmt.Println("sqrt(",b,") : ",Sqrt(b))
  fmt.Println("sqrt(",c,") : ",Sqrt(c))
  fmt.Println("sqrt(",d,") : ",Sqrt(d))
  fmt.Println("sqrt(",e,") : ",Sqrt(e))
  fmt.Println("sqrt(",f,") : ",Sqrt(f))
}

