package main

import "fmt"

func stackingDefers() {
  // deferred function calls are pushed onto a stack
  // when a function returns, its deferred calls are executed in LIFO order
  fmt.Println("counting")

  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }

  fmt.Println("done")
}

func main() {
  // defer deferes the execution of a function until the surrounding function returns
  // the deferred call's arguments are evaluated immediately but the function is not executed until the surrounding function returns
  defer fmt.Println("world")

  fmt.Println("hello")


  stackingDefers()
}
