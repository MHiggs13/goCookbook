package main

import (
  "fmt"
  "runtime"
  "time"
)

func switchNoCondition() {
  t := time.Now()
  // switch with no condition is the same as switch true
  // conditions are at each case
  switch {
  case t.Hour() < 12:
    fmt.Println("Good Morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}

func main() {
  // switch statements are a shorter way to write a swquence of if - else statements
  // it runs the first case whose value is equal to the condition expression
  // Go only runs the selected case, not all the cases taht follow (i.e. break statement is provided automatically)
  // switch cases don't need to be constants
  // values involved don't need to be integers

  fmt.Print("Go runs on ")
  // initial short statement to begin switch
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd, plan9, windows
    fmt.Printf("%s", os)
  }
  switchNoCondition()

}
