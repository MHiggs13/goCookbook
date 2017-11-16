package main

// arg1 is an int, arg2 is a string and foo returns an int
func foo(arg1 int, arg2 string) int  {

}

// we don't need to name the variables
// foo2 receives an int, a string and returns an int
func foo2(int, string) int {
  // declaration of a function variable
  // f is a function that receives a function and an int and returns a function
  f func(func(int,int) int, int) func(int,int) int
}

func basicTypes() {
  bool

  string

  int int8 int16 int32 int64
  uint uint8 uint16 uint32 uint64

  byte  // alias for uint8

  rune  // alias for int32
        // represents a unicode code point

  float32 float64

  complex64 complex128

  // when you need an int you should use int (32 bits wide on 32 bit systems and 64 bits wide on 64 bit systems)
  // unless you need a specific size or an unsigned integer type
}

func typeConversions() {
  // T(v) converts teh value v to the type T
  var i int = 42
  var f float54 = float64(i)
  var u uint = uint(f)
}

func typeInference() {
  // declaring a variable without giving an explicit type, the variable's type is inferred from the value on the right
  var i int
  j := i // j is an int

  // when the right hand side declares a numeric constant, the new variable may be int, float64 or complex128 depending on the precision
  i := 42 // int
  f := 3.142  // float
  g := 0.867 + 0.5i // complex128

}



func main() {

  // types are declared totally differently to how values are assigned
  // this makes it easier to distinguish what is happening
  x int   // x is an int
  p *int    // p is a pointer to an int
  a [3]int    // a is an int array of size 3

  // the distinction between type and expression syntax makes it easy to write and invoke closures
  // TODO: CLOSURE: 
  sum := func(a,b int) int { return a+b} (3,4)
}
