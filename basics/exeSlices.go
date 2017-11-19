package main

import ("code.google.com/p/go-tour/pic")



func Pic(dx, dy int) (a [][]uint8) {
  a = make([][]uint8, dy)

  for i := range a {
    a[i] = make([]uint8, dx)
  }

  for i := range a {
    for j := range a[i] {
      switch {
        case j % 15 == 0:
          a[i][j] = 240
        case j % 3 == 0:
          a[i][j] = 120
        case j % 5 == 0:
          a[i][j] = 150
        default:
          a[i][j] = 100
      }
    }
  }

  return
}

func main() {
  // implement Pic. it should return a slice of length dy
  // each element of which is  aslice of dx 8-bit unsigned integers.
  // When you run the program it will display your picture, interpreting the integers as grayscale (well, bluescale) values
  // the choice of image is up to you. Interesting functions include (x+y)/2, x*y and x^y
  pic.Show(Pic)
}
