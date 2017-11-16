package main

import "fmt"

type Vertex struct {
  // a struct is a collection of fields
  X int
  Y int
}

func accessingFields() {
  // struct fields are accessed using a dot
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.Y)
}

func pointersToStructs() {
  // struct fields can be accessed through a struct pointer
  // to access  field X of a struct when we have a struct pointer "p"
  // we could write (*p).X
  // however we can just write p.X instead, without the expicit difference
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  fmt.Println(v)
}

var (
  // a struct literal denotes a newly allocated struct value by listing the values of its fields
  v1 = Vertex{1, 2}   // has type Vertex
  v2 = Vertex{X: 1}   // Y:0 is implicit
  v3 = Vertex{}       // x:0 and Y:0
  p = &Vertex{1, 2}   // has type *Vertex
)

func structLiterals() {
  // a struct literal denotes a newly allocated struct value by listing the values of its fields
  // a literal is a fixed value i.e 3, 4, 44
  // i.e. creating the struct and putting the values directly in, instead of assinging them later
  fmt.Println(v1, p, v2, v3)
}

func main() {
  fmt.Println(Vertex{1, 2})

  pointersToStructs()
  structLiterals()

}

