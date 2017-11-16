package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  // if statements are like its for loops, the expression need not be surrounded by () but {} are required
  if x < 0 {
          return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func ifShort(x, n, lim float64) float64 {
  // like a for statement the if statement can start witha short statement to execute before the condition
  // variables declared by the statement are only in scope until the end of the if
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}

func ifElse(x, n, lim float64) float64 {
  // variables declared inside an if are also available in subsequent else blocks
  if v:= math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }

}

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
  fmt.Println(ifShort(3, 2, 10), ifShort(3, 3, 20))
  fmt.Println(ifElse(3, 2, 10), ifElse(3, 3, 20))
}
