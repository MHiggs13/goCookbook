package main

import (
  "fmt"
  "io"
  "strings"
)

func main() {
  // the io package specifies the io.Reader interface, which represents the read end of a stream of data
  // the go standard library contains many implementations of these interfaces
  // including files, network connections, compressors, ciphers and others
  // the io.Reader interface has a Read method
  // func (T) Read(b []byte) (n int, err error)
  // Read populates the given byte slice with data and returns the number of bytes populated and an error value
  // it returns an io.EOF error when the stream ends
  r := strings.NewReader("Hello, Reader!")

  b := make([]byte, 8)
  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])
    if err ==  io.EOF {
      break
    }
  }
}
