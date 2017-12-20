package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

func Walk(t *tree.Tree, ch chan int) {
  _walk(t, ch)
  close(ch)
}

// Walk walks the tree t sending all values
// from the tree to channel ch.
func _walk(t *tree.Tree, ch chan int) {
  if t != nil {
    _walk(t.Left, ch)
    ch <- t.Value
    _walk(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)

  for v := range ch1 {
    if v != <-ch2 {
      return false
    }
  }
  return true


}

func main() {
  // Implement Walk
  // Test Walk
  // Implement Same using Walk to determine whether t1 and t2 store the same values
  // Test Same

  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for v := range ch {
    fmt.Println(v)
  }

  fmt.Println(Same(tree.New(1), tree.New(2)))
  fmt.Println(Same(tree.New(1), tree.New(1)))
  fmt.Println(Same(tree.New(2), tree.New(2)))

}
