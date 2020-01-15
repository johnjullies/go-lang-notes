package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
  n, err := reader.r.Read(b)
  for i := range b {
    b[i] = rot13(b[i])
  }
  return n, err
}

func rot13(x byte) byte {
    capital := x >= 'A' && x <= 'Z'
    if !capital && (x < 'a' || x > 'z') {
      return x
    } else if capital && x < 'N' || !capital && x < 'n' {
      x += 13
    } else {
      x -= 13
    }
    return x
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
