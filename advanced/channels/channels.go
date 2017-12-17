package main

import "fmt"

func bufferedChannels () {
  // channels can be buffered, provide the buffer length as the second argument to make to initialize a buffered channel
  // ch := make(chan int, 100)
  // sends to a vuffered channel blcok only when the buffer is full.
  // receives block when the buffer is empty

  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  // ch <- 3  // goroutines are already shut off at this point because the buffer is full
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  //fmt.Println(<-ch)


}

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s{
    sum += v
  }
  c <- sum //send sum to c
}

func main() {
  // channels are a typed conduit through which you can send a nd receive values with the channel operator, <-
  // ch <- v    // send v to channel ch
  // v := <- ch   // receive from ch, and assign value to v
  // the data flows in the direction of the arrow

  // like maps and slices, channels must be created before use:
  // ch := make(chan int)

  // by default, send and receives block until the other side is ready, this allows goroutines to sync,
  // without explicit locks or conditiion variables
  // the example code sums the numbers in a slice, distributing the work between two goroutines
  // once both goroutines have completed their computation, it calculates teh final result.
  s := []int{7, 2, 8, -9, 4, 0}

  c := make(chan int)
  go sum(s[:len(s)/2], c)   // first half
  go sum(s[len(s)/2:], c)   // second half

  // waits on sends to be completed
  x, y := <-c, <-c // receive from c

  fmt.Println(x, y, x+y)

  bufferedChannels()
}
