# More types: structs, slices, and maps
https://tour.golang.org/moretypes/1

## Pointers
Pointers hold the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is nil. 

The `&` operator generates a pointer (memory address) to its operand.  
The `*` operator denotes the pointer's underlying value. 

```go
i := 42
p = &i
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" or "indirecting". 

Unlike C, Go has no pointer arithmetic. More info on the importance of no pointer arithmetic: https://stackoverflow.com/questions/32700999/pointer-arithmetic-in-go

## Struct
Collection of fields. Similar to Python's `dict`

```go
type Vertex struct {
  X int
  Y int
}

func main() {
  fmt.Println(Vertex{1, 2})
}
```

Fields are accessed using dot notation

```go
v := Vertex{1, 2}
fmt.Println(v)    // {1 2}
v.X = 4
fmt.Println(v)    // {4 2}
```

### Pointers to struct
To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X` or just `p.X`

### Struct literals
```go
v1 = Vertex{1, 2}  // has type Vertex
v2 = Vertex{X: 1}  // Y:0 is implicit
v3 = Vertex{}      // X:0 and Y:0
p  = &Vertex{1, 2} // has type *Vertex
```

## Arrays
```go
var a [10]int // declares a variable `a` as an array of ten integers
```

Array's length is part of its _type_ so it cannot be resized (i.e. fixed size)

## Slices
Dynamically-sized, flexible view into the elements of an array. Much common in practice than arrays.

```go
primes := [6]int{2, 3, 5, 7, 11, 13}

var s []int = primes[1:4]
fmt.Println(s)  // [3 5 7]
```

### Slices are like references to arrays
A slice does not store any data, just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes. 

```go
names := [4]string{
  "John",
  "Paul",
  "George",
  "Ringo",
}
fmt.Println(names)  // [John Paul George Ringo]

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)   // [John Paul] [Paul George]

b[0] = "XXX"
fmt.Println(a, b)   // [John XXX] [XXX George]
fmt.Println(names)  // [John XXX George Ringo]
```

### Slice literals
Like an array without the length. Also builds a slice that references it.

```go
r := []int{2, 3, 5, 7, 11, 13}
fmt.Println(r)
fmt.Println(r[0:1])

// [2 3 5 7 11 13]
// [2]
```

### Slice defaults
Low and high bounds can be omitted.

`0` default of low bound  
`length` of the slice for the high bound

### Slice length and capacity
A slice has both length and capacity.
- length is the number of elements in the slice; `len(s)`
- capacity is the number of elements in the referenced array; `cap(s)`

```go
func main() {
  s := []int{2, 3, 5, 7, 11, 13}
  printSlice(s)

  // Slice the slice to give it zero length.
  s = s[:0]
  printSlice(s)

  // Extend its length.
  s = s[:4]
  printSlice(s)

  // Drop its first two values.
  s = s[2:]
  printSlice(s)
  
  // Extend it again
  s = s[:4]
  printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// len=6 cap=6 [2 3 5 7 11 13]
// len=0 cap=6 []
// len=4 cap=6 [2 3 5 7]
// len=2 cap=4 [5 7]
// len=4 cap=4 [5 7 11 13]
```

### Nil slices
Zero value of a slice is `nil`

### The built-in `make` function
Slices can be created with `make`. This is how you make dynamically sized arrays.

```go
a := make([]int, 5)
printSlice("a", a)

b := make([]int, 0, 5)
printSlice("b", b)

// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
```


### Slices of slices
Slices can contain any type, including other slices.

```go
import (
  "fmt"
  "strings"
)

func main() {
  // Create a tic-tac-toe board.
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // The players take turns.
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
}
```

### Appending to a slice
```go
var s []int
printSlice(s)
s = append(s, 2, 3, 4)
printSlice(s)

// len=0 cap=0 []
// len=3 cap=4 [2 3 4]
```

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array. 

Further reading: https://blog.golang.org/go-slices-usage-and-internals

## Range
Used in `for` loop.  
Iterates over a slice or map.

Returns two values when ranging over a slice, the index and a copy of the element in that index.

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}

// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 2**4 = 16
// 2**5 = 32
// 2**6 = 64
// 2**7 = 128
```

You can skip the index or value by assigning to `_`.

```go
for i, _ := range pow
for _, value := range pow
```

If you only want the index, you can omit the second variable.

```go
for i := range pow
```
