# Testing
Go has a built-in lightweight test framework

Create a file ending in `_test.go` that contains functions named `TestXXX` with signature `func (t *testing.T)`

```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    {"Hello, world", "dlrow ,olleH"},
    {"Hello, 世界", "界世 ,olleH"},
    {"", ""},
  }
  for _, c := range cases {
    got := ReverseRunes(c.in)
    if got != c.want {
      t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
    }
  }
}
```

then run the test with `go test`
