 package main

 import (
  "fmt"
  "time"
 )

 func main() {
  // the default case in a select is used if no other ase is ready
  // use a default case to try send or receive without blocking
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)
  for {
    select {
      case <-tick:
        fmt.Println("tick.")
      case <-boom:
        fmt.Println("BOOM!")
        return
      default:
        // not receiving anything so print this
        fmt.Println("    .")
        time.Sleep(50 * time.Millisecond)

    }
  }
}
