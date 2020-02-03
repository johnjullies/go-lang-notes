# Go Basics

## Packages
Every Go program is made up of packages. Starts with the package declaration.

```go
package main
```

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`.

A package is a collection of source files in the same directory. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

A **module** is a collection of related Go packages that are released together. A **repository** is a collection of one ore more modules.

## Imports
Parenthesized, "factored" import statement. 

```go
import (
  "fmt"
  "math"
)
```

When importing a package, you can refer only to its exported names. Exported names must start with a capital letter. Any "unexported" names are not accessible from outside the package. 

```go
fmt.Println(math.Pi)
```

## Functions
The type of arguments comes _after_ the variable name.

```go
func add(x int, y int) int {
  return x + y
}
```

Article on Go's declaration syntax: https://blog.golang.org/gos-declaration-syntax

Function parameters type sharing

```go
func add(x, y int) int {
  return x + y
}
```

Functions can return _any_ number of results.

```go
func swap(x, y string) (string, string) {
  return y, x
}
```

### Named return values
Return values can be named at the top of the function

```go
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
```

"naked" return = return statement without arguments. Returns the named values. Should __only__ be used in short functions.

## Variables
`var` statement declares list of variables. Declaration syntax same with functions.

`var` statement can be package (global) or function level.

```go
var c, python, java bool

func main() {
  var i int
  fmt.Println(i, c, python, java) // 0 false false false
}
```

### Variables with initializers
`var` declaration can have initializers (assignment operator), one per variable. The type can be omitted. The variable will take the type of the initializer.

```go
var i, j int = 1, 2 // type is optional

func main() {
  var c, python, java = true, false, "no!" // type is omitted
}
```

### Short variable declarations
Can __only__ be used inside a function. Can be used in place of `var` declaration with implicit type.

Every statement must begin with a keyword outside a function so the `:=` construct is not available.

```go
func main() {
  var i, j int = 1, 2
  k := 3

  fmt.Println(i, j, k)
}
```

### Basic Types
```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

`var` declarations can be factored like `import` statements

```go
import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe    bool        = false
  MaxInt  uint64      = 1<<64 - 1
  z       complex128  = cmplx.
)
```

Use `int` for integers unless specific reason to use sized or unsigned integer type.

### Zero values
Variables without initializers are given _zero value_

The zero value is:
- 0 for numeric types
- false for bool
- "" for string

### Type conversions
`T(v)` converts the value `v` to type `T`

Explicit conversion required for assignment bet. different types

```go
var x, y int = 3, 4
var f float64 = math.Sqrt(float64(x*x + y*y))
var z uint = uint(f)
```

### Type inference
The type of the __right__ hand side of a variable declaration _without_ explicit type is the type of the variable.

```go
var i int
j := i // j is int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

## Constants
Declared like variables but using the `const` keyword. Can also be factored.

Cannot be declared using the `:=` syntax

`const Pi = 3.14`

### Numeric constants
High precision values.

Takes the type needed by its context.
