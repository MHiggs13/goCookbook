package main

import "fmt"

// Go only has one looping construct, the for loop
// For loop has 3 components, separated by semicolons
      // init statement: executed before the first iteration
      // condition statement: evaluated before every iteration
      // post statement: executed at the end of every iteration

func optionalInitAndPost() {
  // the init and post statement are optional
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}

func gosWhiles() {
  // can drop the semi colons to make a while called for
  sum := 1
  for sum < 1000 {
    sum += sum
  }
}

func infiniteLoop() {
  // omitting the loop condition loops it forever
  for {
  }
}


func main() {
  sum := 0
  for i := 0;i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
