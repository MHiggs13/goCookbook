package main

import "fmt"

func add(x int, y int) int {
  // func add takes two parameters of type int and returns an int
  return x + y
}

func add2(x, y int) int {
  // for consecutive named function parameters the type can be omitted from all but the last parameter
  return x + y
}

func swap(x, y string) (string, string) {
  // a function can return a number of results        
  return y, x
}

func split(sum int) (x, y int) {
  // contains named return values
  // naked returns should only be used in short functions as readability can be harmed in longer functions
  x = sum * 4/9
  y = sum - x
  return    // naked return - returns the named return values (x and y both ints) 
}

func main() {
  fmt.Println(add(52, 13))
}

