package main

import (
  "fmt"
  "sync"
  "time"
)

// SafeCounter is safe to use concurrently
type SafeCounter struct {
  v map[string]int
  mux sync.Mutex
}

// Inc increments the counter for the fiven key.
func (c *SafeCounter) Inc(key string) {
  c.mux.Lock()
  // Lock so only one goroutine at a time can access the map c.v.
  c.v[key]++
  c.mux.Unlock()
}

// Value returns the current value of the counter for the fiven key
func (c *SafeCounter) Value(key string) int {
  c.mux.Lock()
  // Lock so only one goroutine at a time can access the map c.v.
  defer c.mux.Unlock()    // defers until the surrounding function returns
  return c.v[key]
}

func main() {
  // c is of type SafeCounter which provides mutex (sync.Mutex) for v (a map[string]int)
  c := SafeCounter{v: make(map[string]int)}
  for i := 0; i < 1000; i++ {
    go c.Inc("somekey")
  }

  time.Sleep(time.Second)
  fmt.Println(c.Value("somekey"))
}
