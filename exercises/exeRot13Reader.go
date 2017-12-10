package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error){
  n, err = rot.r.Read(p)
  for i := 0; i < len(p); i++ {
    if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
      p[i] += 13
    } else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
      p[i] -= 13
    }
  }
  return
}

func main() {
  //  a common pattern is an io.Reader that wraps another io.Reader, modigying the stream
  // the gzip.NewReader function takes an io.Reader (a  stream of compressed data) and returns a *gzip.Reader
  // this also implements io.Reader (a stream of compressed data)
  // implement a rot13Reader that implements io.Reader and reads form an io.Reader,
  // modifying the stream by applying the rot13 substitution cipher to all alphabetical characters
  // implement it's Read method
  s := strings.NewReader("lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
