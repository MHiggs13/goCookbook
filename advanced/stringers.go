package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  // Stringer is defined by the fmt package
  // type String interface {
  //   String() string
  // }
  // a Stringer is a type that can descrive itself as a string
  // the fmt package (and many others) look for this interface to print values
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphood Beeblebrox", 9001}
  fmt.Println(a, z)
}
