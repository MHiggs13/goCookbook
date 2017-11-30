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

func functionValues(fn func(int, int) int) int {
  // functions are values too
  // they can be passed around just like other values
  // function values may be used as function arguements and return values
  return fn(10, 10)
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
      sum += x
      return sum
    }
}

func functionClosures() {
  // functions may be closures
  // closure: function value that references variables from outside its body
  // the function may access and assign to the referenced variables
  // in this sense the function is "bound" to the variables
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
      )
  }
}

func main() {
  add3 := func(x, y int) int {
    return x + y
  }
  fmt.Println(add(52, 13))
  fmt.Println(functionValues(add3))
  fmt.Println()
  functionClosures()
}

