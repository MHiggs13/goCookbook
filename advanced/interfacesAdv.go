package main

import (
  "fmt"
  "math"
)

type Abser1 interface {
  // an interface type is defined a s a set of method signatures
  // a value of interface type can hold any value that implements those methods
  Abs1() float64
}

type MyFloat float64

func (f MyFloat) Abs1() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(-f)
}


type Abser2 interface {
  // an interface type is defined a s a set of method signatures
  // a value of interface type can hold any value that implements those methods
  Abs2() float64
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs2() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


type I interface {
  // under the covers, interface values can be though of as a tuple of a value and a concrete type: (value, type)
  // an interface value holds a value of a specific underlying concrete type
  // calling a method on an interface value executes the method of the same name on its underlying type
  M()
}

type T struct {
  S string
}

// this method type T implements the interface I,
// but we don't need to explicitly declare that it does so
func (t *T) M() {
  // if the concrete value inside the interface itself is nil, the method will be called with a nil receiver
  // in some languages this would trigger a null pointer exception, but in Go
  // it is common ot write methods that gracefully handle being called with a nil receiver
  // note that an interface value that holds a nil concrete value is itself non-nil
  if t == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(t.S)
}

func interfacesImplicit() {
  // a type implements an interface by implementing its methods
  // implicit interfaces, decouple the definition of an interface from its implementation
  // which could appear in any package without prearrangement...
  var i I = &T{"hello"}
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func nilInterface() {
  // a nil interface value holds neither value nor concrete type
  // calling a method on a nil interface is aruntime error 
  // because there is no type inside the interface typle to indicate which concrete method to call
  var i I
  describe(i)
  i.M()
}

func describe2(i interface{})  {
  fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
  // the interface type taht specifies zero methods is known as the empty interface
  var i interface{}
  describe2(i)

  // an empty interface may hold values of any type, every type implements at least zero methods
  // empty interfaces are used by code taht handle values of unknown type
  // fo rexample, fmt.Print takes any number of arguements of type interface{}
  i = 42
  describe2(i)

  i = "hello"
  describe2(i)
}


func main() {
  var a1 Abser1
  var a2 Abser2
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a1 = f // a Myfloat implements Abser
  a2 = &v // a *Vertex implements Avser

  // v is a Vertex (not *Vertex) and does NOT implement Abser
  //a3 = v

  fmt.Println(a1.Abs1())
  fmt.Println(a2.Abs2())
  // error
  //fmt.Println(a3.Abs2())

  interfacesImplicit()
  //nilInterface()
  emptyInterface()
}
