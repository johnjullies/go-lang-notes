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
