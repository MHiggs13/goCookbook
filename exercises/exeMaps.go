package main

import (
  //"fmt"
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  arrS:=strings.Fields(s)
  var m = make(map[string]int)
  for _, v := range arrS {
    // todo: doesn't work for "a" as it counts occurences of a within words such as "apple" and not a on it's own 
    m[v] = strings.Count(s, v)
  }
  return m
}

func main() {
  // implement WordCount
  // it should return a map of the count of each word in the string s
  // the wc.Test function runs a test suite against the provided function and prints success or failure
  // strings.Fields might be useful
  //fmt.Println(WordCount("this is a test"))
  wc.Test(WordCount)
}
