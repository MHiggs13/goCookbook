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

func rangeFor() {
  // the range form of the for loop iterates over a slice or a map
  // when ranging over a slice, two values are returned for each iteration. 
  // the first is the index, and the second is a copy of the element at that index
  var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

  // two values - i: index, v:value
  for i,v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }

  // you can skip the index of value by assigning to _
  // if you only want the index, drop the ", value" entirely

  // skip index
  for _, value := range pow {
    fmt.Printf("%d\n", value)
  }

  // skip value
  for i := range pow {
    pow[i] = 1 << uint(i)   // == 2**1
  }

}


func main() {
  sum := 0
  for i := 0;i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
