package main

import (
  "fmt"
  )

func main() {
  // Go has pointers - a pointer holds the memory address of a value

  // the type *T is a pointer to a T value
  // it's zero value is nil

  i, j := 42, 2701

  // the & command generates a pointer to it's operand (points to i)
  p := &i
  // read what is at the pointer p (i is read through the pointer)
  fmt.Println(*p)
  // set the value at which the pointer p points at (set the value of i
  *p = 21
  // print out the new value of i
  fmt.Println(i)


  fmt.Println("j: ", j)
  pj := &j
  fmt.Println("pj: ", pj)
  fmt.Println("*pj: ", *pj)

  *pj = 2112
  fmt.Println("*pj = 2112")


  fmt.Println("j: ", j)
  fmt.Println("pj: ", pj)
  fmt.Println("*pj: ", *pj)



}
