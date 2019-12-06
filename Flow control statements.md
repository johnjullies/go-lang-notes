# Flow control statements

## For
The only looping in Go is `for` loop.

Familiar three parts seperated by semicolons _without_ parenthesis. Brackets are __always__ required.

```go
for i := 0; i < 10; i++ {
  // ...do until i < 10
}
```

The init and post statements are optional so are the semicolons. For is Go's `while`.

```go
sum := 1
for sum < 1000 {
  sum += sum
}
fmt.Println(sum) // 1024
```

## If
Like `for`, no `()` needed but `{}` are required.

```go
if 2 > 0 {
  return "Positive"
}
```

Like `for`, can start with short init statement. Variables declared here are in scope until the end of the `if` and inside any `else` block.

```go
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    fmt.Println(v)
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func main() {
  pow(3, 2, 10) // 9
  pow(3, 3, 20) // 27 >= 20
}
```

## Switch
Familiar syntax except no `break` statement needed. Also accepts init statement.

```go
fmt.Print("Go runs on ")
switch os := runtime.GOOS; os {
case "darwin":
  fmt.Println("OS X.")
case "linux":
  fmt.Println("Linux.")
default:
  // freebsd, openbsd,
  // plan9, windows...
  fmt.Printf("%s.\n", os)
}
```

Can be written without a condition. Same as `switch true`

```go
switch {
  case conditionA:
    //
  case conditionB:
    //
  default:
    //
}
```

## Defer
Delays execution of function until the surrounding function returns. The deferred call's arguments are evaluated immediately but not executed until the surrounding function returns.

```go
func main() {
  defer fmt.Println("world")

  fmt.Println("hello")
}

// prints hello world
```

### Stacking defers
All deferred function calls are pushed to a stack. When the function returns, its deferred calls are executed in last-in-first-out order.

Learn more: https://blog.golang.org/defer-panic-and-recover
