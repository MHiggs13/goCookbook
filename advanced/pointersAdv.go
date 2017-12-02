package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
  // method defined on *Vertex
  // Methods with pointer receivers can modify the value to which the receiver points
  // since methods often need to modify their receiver, pointer receivers are more common than value reeivers
  // if Scale worked on a value receiver, the method would operate on a copy of the original Vertex value
  // this is the same behaviour as any other function argument
  // The scale method must have  apointer receiver to change the Vertex value declared in main
  v.X = v.X * f
  v.Y = v.Y * f
}

func pointerReceivers () {
  // you can declare methods with pointer receivers
  // this means the receiver type ahs the literal syntax *T for type T
  // T itself cannot be a pointer (i.e. can't be *MyInt)
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
  pointerFunction(&v)
  fmt.Println(v.Abs())
}

func pointerFunction(v *Vertex) {
  // to change the actual values, a pointer needs to be passed in
  v.X = 10
  v.Y = 10
}

func choosingValueOrPointerReceiver ()  {
  // there are two reasons to use a pointer
  // 1.) so that the method can modify the value that its receiver poitns to
  // 2.) avoid copying the value on each method call, can be more efficient if teh receiver is a large struct

  // in general, all methods on a given type should have either value or pointer receivers, not a mixture of both
}

func main() {
  pointerReceivers()
  
  // functions with a pointer argument must take a pointer
  // methods with a pointer receiver can take either a value or a pointer 

  // the equivalent thing happens in the reverse direction

  // a function that takes a  value argument must take a value of that specific type
  // methods with a value receiver can take either a value or a pointer
}
