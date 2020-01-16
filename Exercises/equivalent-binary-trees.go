package main

import (
  "golang.org/x/tour/tree"
  "fmt"
  "time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  if t == nil {
    return
  } else if t.Left == nil {
    ch <- t.Value
    if t.Right != nil {
      Walk(t.Right, ch)
    }
    return
  } else {
    Walk(t.Left, ch)
  }
  ch <- t.Value
  if t.Right != nil {
    Walk(t.Right, ch)
  }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  var result bool
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  go func() {
    for i := range ch1 {
      if i == <-ch2 {
        result = true
      } else {
        result = false
      }
    }
  }()
  time.Sleep(100 * time.Millisecond)
  return result
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  go func() {
    fmt.Printf("Test:")
    for i := range ch {
      fmt.Printf("%v,", i)
    }
  }()
  t1 := tree.New(1)
  t2 := tree.New(1)
  fmt.Println(Same(t1, t2))
  time.Sleep(1 * time.Second)
}
