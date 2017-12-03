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

func typeAssertions() {
  // type assertions provide access to an interface value's underlying concrete value
  // t := i.(T)
  // this statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t
  // if i does not hold a T, the statement will trigger a panic
  // to test whether an interface value holds a specific type a type assertion can return two values
  // the underlying balue and a boolean value that reports whether the assertion succeeded
  // t, ok := i.(T)
  // if i holds a T, then t will be the underlying balue and ok will be true
  // if not, ok will be false and t will be the zero value of type T and no panic occurs
  // note the similarity between this suntax and that of reading from a map
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  //f := i.(float64)  // PANIC
  //fmt.Println(f)
}

func typeSwitches(i interface{}) {
  // type switches are a construct that permits several type assertions in series
  // a type switch is like a regular switch statement, but the case in a type switch specify types(not values)
  // those values are compared against the type of the value held by the given interface value
  // the declaration in a type switch has the same syntax as a type assertion, but the specific type T is replaced with the keyword type
  switch v := i.(type) {
    case int:
      fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
      fmt.Printf("%q is %v bytes long\n", v, len(v))  // variable v is type string
    default:
      fmt.Printf("I don't know about type %T!\n", v)
  }
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

  typeAssertions()

  typeSwitches("hello")
}
