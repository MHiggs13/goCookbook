package main

import (
  "fmt"
)

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i< n; i++ {
    c <- x
    x, y = y, x+y
  }
  // close channel to indicate that no more values will be sent
  close(c)
}

func main() {
  // a sender can close a channel to indicate that no more values will be sent
  // receivers can test whether a channel has been closed by assigning a second parameter to receiver expression
  // v, ok := <-ch
  // if ok is false there are no more values to receive and channel is closed
  // the loop for i := range c receives values from the channel repeatedly until it is closed
  // only the sender should close a channel, never the receiver - sending on a closed channel will cause panic
  // channels aren't like files, you don't usually need to close them
  // Closing them is only necessary when the receiver mst be told there are no more values coming, such as a range loop

  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  // receives values until channel is closed
  for i := range c {
    fmt.Println(i)
  }
}
