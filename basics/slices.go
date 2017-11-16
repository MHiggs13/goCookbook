package main

import (
  "fmt"
  "strings"
)

func slicesReferenceArrays() {
  // a slice does not store any data, it just descrives a section of an underlying array
  // changing the elements of a slice modifies the correspoinding elements of its underlying array
  // other slices that share the same underlying array will see the same changes
  names := [4]string{
    "John",
    "Paul",
    "George",
    "Ringo",
  }

  fmt.Println(names)
  
  // a and b are slices of the names array
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a, b)

  // change the first element of b (i.e. element [1] of names, and element [1] of a)
  b[0] = "XXX"
  fmt.Println(a, b)
  fmt.Println(names)
}

func sliceLiterals() {
  // a slice literal is like an array literal without the length
  // ARRAY literal
  a := [3]bool{true, true, false}

  // SLICE literal - creates the same array as above then builds a slice that references it
  s := []bool{true, true, false}

  fmt.Println(a)
  fmt.Println(s)

  // a slice of structs(structs of type int and bool)
  ss := []struct {
    i int
    b bool
    }{
      {2, true},
      {3, false},
      {5, true},
      {7, true},
      {11, false},
    }
  fmt.Println(ss)
}

func sliceDefaults() {
  // when slicing, the high and low bounds may be ommitted
  // default is zero for LOW bounds
  // default is length of slice for HIGH bounds
  var a [10]int
  
  a[0] = 0
  a[1] = 10
  a[2] = 20
  a[3] = 30
  a[4] = 40
  a[5] = 50
  a[6] = 60
  a[7] = 70
  a[8] = 80
  a[9] = 90

  // these are all equivalent
  printSlice(a[0:10])
  printSlice(a[:10])
  printSlice(a[0:])
  printSlice(a[:])
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printMakeSlice(s string, x []int) {
  fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}

func sliceLengthAndCapacity() {
  // a slice has a length and a capacity
  // length: number of elements in the slice
  // capacity: number of elements in the underlying array counting from the first element in the slice
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // slices to give slice 0 length
  s = s[:0]
  printSlice(s)

  // extend it's length (as it has enough capacity)
  s = s[:4]
  printSlice(s)

  // drop its first two values
  s=s[2:]
  printSlice(s)
}

func nilSlices() {
  // the zero value of a slice is nil
  // a nil slice has  alenght and capacity of 0 and has no underlying array
  var s []int
  fmt.Println(s, len(s), cap(s)) 
  if s == nil {
    fmt.Println("nil!")
  }
}

func creatingSliceWithMake() {
  // slices can be created with the built-in make function
  // this is how you create dynamically-sized arrays
  // the make function allocates a zeroed array and returns a slice that referes to that array
  //
  a := make([]int, 5) // len(a) = 5
  printMakeSlice("a", a)

  b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
  printMakeSlice("b", b)

  c := b[:2] 
  printMakeSlice("c", c)

  d := c[2:5]
  printMakeSlice("d", d)
}

func sliceOfSlices() {
  // a slice can contain any type, including other slices
  // create a tic-tac-toe board
  // is a slice containing slices (2d slice)
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // the player takes turns
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }

}

func appendToSlice() {
  // it is common to append elements to a slice
  // Go provides a built in append function

  // documentation descrives append as 
  // func append(s []T, vs ...T) []T
  // s is a slice of type T
  // rest of parameters are T values to append
  // returned value is of a slice containing all the original slice elements along with the provided values

  // if the backing array of s is too small to fit all the given values, a bigger array will be allocated
  // the returned slice will point to a newly allocated array

  // slice of type int
  var s []int
  printSlice(s)

  // append Works on nil slices
  s = append(s, 0)
  printSlice(s)

  // slice grows as needed
  s = append(s, 1)
  printSlice(s)

  // we can add more than one element at a time
  s = append(s, 2, 3, 4)
  printSlice(s)
}

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

  slicesReferenceArrays()
  sliceLiterals()
  sliceLengthAndCapacity()
  nilSlices()
  creatingSliceWithMake()
  sliceOfSlices()
  appendToSlice()
}
