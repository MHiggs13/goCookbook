package main

import (
  "fmt"
  "time"
)

func say(s string) {
  for i := 0; i < 5; i++ {
    time.Sleep(100 * time.Millisecond)
    fmt.Println(s)
  }
}

func main() {
  // a goroutine is a ightweight thread managed by the Go runtime
  // go f(x, y, z)
  // starts a new goroutine running
  // the evaluation of f,x,y,z happens in the current goroutine and the execution of f happens in the new goroutine
  // goroutines run in the same address space, so access shared memory myst be synchronized
  // the sync pacakge provides useful primitives, although you won't need them much in go as there are other primitives
  go say("world")
  say("hello")
}
