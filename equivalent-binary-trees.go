package main

import (
  "golang.org/x/tour/tree"
  "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  if t == nil { return }
  Walk(t.Left, ch)
  ch <- t.Value
  Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1, ch2 := make(chan int), make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for i := 0; i < 10; i++ {
    val1, val2 := <-ch1, <-ch2
    if val1 != val2 {
      return false
    }
  }
  return true
}

func main() {
  // test for Walk
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for i := 0; i < 10; i++ {
    fmt.Println(<-ch)
  }

  // test for Same
  fmt.Println(Same(tree.New(1), tree.New(1))) // true
  fmt.Println(Same(tree.New(1), tree.New(2))) // false
}
