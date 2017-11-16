package main

import "fmt"

// variable at package level
// var declares a list of variables. variable c, python and java are bools
var c, python, java bool

// INITIALIZER
// var declaration with an initializer fo reach variable
var i2, j2 int = 1, 2

func main() {
  // variable at function level
  // variable i is an int
  var i int
  fmt.Println(i, c, python, java)

  // INITIALIZER
  // initializer is present, type can be omitted, in this case type of initializer is used
  var c2, python2, java2 = true, false, "no!"
  fmt.Println(i2, j2, c2, python2, java2)

  // SHORT ASSIGNMENET
  // inside a function the short assignment statement can be used in place of a var declaration
  // outside a function every statement begins with a keyword
  k := 3
  // no need to use var
  test, moo := true, "no!"

  // ZERO VALUES
  // variables declared without an explicit initial value are fiven their zero value
  var i int // 0
  var f float64 // 0
  var b bool  // false
  var s string  // ""
}
