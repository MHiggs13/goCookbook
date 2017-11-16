package main

import "fmt"

func main() {
  // the type [n]T is an array of n values of type T
  
  // an array, a, with 2 values of type string
  var a [2] string
  a[0] = "Hello"    // set first value in a equal to "Hello"
  a[1] = "World"
  fmt.Println(a[0], a[1])
  fmt.Println(a)

  // an array of 6 integers - size is fixed
  primes := [6]int{2, 3, 5, 7, 11, 13}
  fmt.Println(primes)
}
