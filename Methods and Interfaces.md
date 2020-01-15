# Methods and interfaces
https://tour.golang.org/methods/1

## Methods
Functions defined _on_ types with a _special_ receiver argument

```go
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// in this example, the Abs method
// has a receiver of type Vertex named v
```

### Methods are `func`s
A Go method is just a function with a receiver argument.

```go
// Here, Abs is written as a regular function
// with no change in functionality

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

### Methods can be on non-struct types
```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
  // ...
}
```

Methods can only be declared with a receiver whose **type** is defined in the **same package** as the method.

### Pointer receivers
```go
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

- the receiver type has literal syntax (using asterisk)
- the type cannot itself be a pointer (such as `*int`)
- can **modify** the value to which the receiver points

### Methods and pointer indirection
Functions with a pointer argument must take a pointer


```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

while methods with pointer receivers take either value or a pointer

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

As a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

It's the same thing in the reverse direction. Functions that take a value argument must take a value of that specific type.

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

### Value or pointer receiver
2 reasons to use a pointer receiver
- so that the method can modify the value
- to avoid copying the value on each method call. Can be efficient on large structs.

## Interfaces
- defined as a set of method signatures
- a value of interface type can hold any value that implements those methods
- implemented **implicitly**. no need for `implements` keyword

### Interface values
Under the hood, interface values can be thought of as a tuple of a value and their concrete type:

```
(value, type)
```

An interface value holds a value of a specific underlying concrete type.

```go
var i I

i = &T{"Hello"}
describe(i)       // (&{Hello}, *main.T)
i.M()             // Hello
```

Calling the method of the interface value executes the method of the same name on its underlying type.

#### nil underlying values
If the concrete value inside the interface is nil, the method will be called with a nil receiver.

In Go, it is common to write methods that gracefully handle being called with a nil receiver.

**Note** that an interface value that holds a nil concrete value is itself non-nil.

#### nil interface values
Holds neither value nor concrete type.

Run-time error because there is no type inside the *interface tuple* to indicate which concrete method to call.

### The empty interface
```go
interface{}
```
- zero methods
- may hold values of any type (every type implements at least zero methods)
- used by code that handles values of unknown type

### Type assertions
```go
t := i.(T)
```
- asserts that `i` is of type `T` and assigns the underlying value to `t`
- if i does not hold a T, panic occurs

Type assertion can return two values

```go
t, ok := i.(T)
```
- if `i` holds a `T`, then `t` will be the underlying value and `ok` will be true
- if not, `ok` will be false, `t` will be the zero value of type `T`, and no panic occurs

### Type switch
Like a regular switch except cases are types not values

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

`type` keyword is used in the declaration

When there is no match, the variable `v` is of the same interface type and value as `i`.

### Stringers
One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.

```go
type Stringer interface {
    String() string
}
```

Stringers are types that can describe itself as a string. Packages look for this interface to print values.

### Errors
Built-in interface like `fmt.Stringer` and packages look for `error` for printing values

```go
type error interface {
    Error() string
}
```

Functions often return an error value, and calling code should handle errors by testing whether the error equals `nil`.

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

`nil` error = success

### Readers
The `io.Reader` interface has a Read method:

```go
func (T) Read(b []byte) (n int, err error)
```

Read populates the given byte slice with data and returns the *number* of bytes populated and an *error* value. It returns an io.EOF error when the stream ends.

The Go standard library contains [many implementations](https://golang.org/search?q=Read#Global) of these interfaces, including files, network connections, compressors, ciphers, and others.

### Images
```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**Note** the `Rectangle` return value of the `Bounds` method is actually an `image.Rectangle`, as the declaration is inside package `image`.

(See [the documentation](https://golang.org/pkg/image/#Image) for all the details.)

The `color.Color` and `color.Model` types are also interfaces, but we'll ignore that by using the predefined implementations `color.RGBA` and `color.RGBAModel`. These interfaces and types are specified by the `image/color` package
