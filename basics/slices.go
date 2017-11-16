package main

import "fmt"

func main() {
  // an array has a fixed size
  // a slice is dynamically sized, generally slices are more common than arrays

  // an array of 6 ints
  primes := [6]int{2, 3, 5, 7, 11, 13}

  // the type []T is a slice with element of type T
  // a slice is created by specifying two indices, a low and high bound, separated by a colon
  // this selects a half-open range which includes the first element but excludes the last one
  // a slice, s, with elements of the type int - from the array primes with values 1,2 and 3
  var s []int = primes[1:4]
  fmt.Println(s)

}
