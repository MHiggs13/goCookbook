package main

import "fmt"

// constants are declared like variables, but witht he const keyword
// constsants can be character, stirng, boolean, or numeric values
// constants cannot be declared using the := syntax
const Pi = 3.14


// numeric constants are high-precision values
// an untyped constant takes the type needed by its context
// an int can store at maximum a 64-bit integer and sometimes less
const (
  // Create a huge number by shifting a 1 bit left 100 places
  // i.e. the binary number that is 1 followed by 100 zeroes
  Big = 1 << 100

  // Shift it right again 99 places, so we end up with 1<<1, or 2
  Small = Big >> 99

)

func needInt(x int) int {
  return x*10+1
}
func needFloat(x float64) float64 {
  return x * 0.1
}





func main() {
  const World = "Dia duit"
  fmt.Println("Hello ", World)
  fmt.Println("Happy ", Pi, " Day.")

  const Truth = true
  fmt.Println("Go rules?", Truth)

  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  // fmt.Println(needInt(Big))  // causes overflow 
  fmt.Println(needFloat(Big))

}
