package main

import (
  "fmt"
  "time"
)

type MyError struct {
  When time.Time
  What string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
  // go programs express error state with error values
  // the error type is a built-in interface similar to fmt.Stringeri
  // type error interface { Error() string }
  // as with fmt.String, the fmt package looks for the rror interface when printing values
  // functions often return an error value, and calling code should handle errors by testing whether the error equals nil
  // i, err := strconv.Atoi("42")
  // if err != nil...
  return &MyError{
    time.Now(),
    "it didn't work",
  }
}

func main() {
  if err := run(); err != nil {
    fmt.Println(err)
  }
}
