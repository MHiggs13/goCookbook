package main

import "fmt"

func fibonacci(c, quit chan int) {
  x, y := 0, 1
  // infinite loop
  for {
    // select statement lets a goroutine wait on multiple communication operations
    // a select blocks until one of its cases can run, then it executes taht case
    // if multiple are ready it chooses one at random
    select {
      case c <- x:
        x, y = y, x+y
      case <-quit:
        fmt.Println("quit")
        return
    }
  }
}

func main() {
  c := make(chan int)
  quit := make(chan int)
  // goroutine, function happens in parallel
  go func() {
    for i := 0; i < 10; i++ {
      fmt.Println(<-c)
    }
    quit <- 0
  }()
  // function called with values sent by go func()
  fibonacci(c, quit)
}
