package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  // this method has a receiver of type Vertex named v          
  // it can be called with v.Abs()
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs2(v Vertex) float64 {
  // here abs is written as a function with no receiver arguement
  // it can be called with Abs2(v) instead
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// loacally declared type for receiver
type MyFloat float64

func (f MyFloat) Abs3() float64 {
  // methods can be declared on non-struct types
  // a method can only be declared witha  receiver whose type is defined in the same package as the method
  // you cannot declare a method with a receiver whose type is defined in another packae (includes built-in types)
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  // go does not have classes, you can define methods on types however
  // a method is a function with a special receiver argument
  // the receiver appears in its own argument list between the func keyword and the method name
  v := Vertex{3, 4}
  // can then use Abs as a "method"
  fmt.Println(v.Abs())

  fmt.Println(Abs2(v))
}
