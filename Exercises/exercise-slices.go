package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  pic := make([][]uint8, dy)
  for i := range pic {
    pic[i] = make([]uint8, dx)
  }
  for y := range pic {
    for x := range pic[y] {
      // pic[x][y] = uint8(  )
      // pic[x][y] = uint8( (x+y)/2 )
      // pic[x][y] = uint8( x*x+y*y )
      pic[x][y] = uint8( x^y )
    }
  }
  return pic
}

func main() {
  pic.Show(Pic)
}
