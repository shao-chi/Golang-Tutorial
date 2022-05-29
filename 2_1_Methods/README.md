Chapter 2-1. Methods

* [Methods](#methods)
    * [Pointer receivers](#pointer-receivers)
    * [Pointers and functions](#pointers-and-functions)
    * [Methods and pointer indirection](#methods-and-pointer-indirection)
    * [Choosing a value or pointer receiver](#choosing-a-value-or-pointer-receiver)

### Methods

> Go does not have classes.

* Define methods on types.
* A method is a **function** with a special **receiver** argument.
* The receiver appears in its own argument list between the `func` keyword and the method name.

```go
type Vertex struct {
	X, Y float64
}

// method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

```go
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

* You can only declare a method with a receiver whose type is defined in the same package as the method.
* You cannot declare a method with a receiver whose type is defined in another package.

```go
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

### Pointer receivers

* The receiver type has the literal syntax `*T` for some type `T`. (Also, `T` cannot itself be a pointer such as `*int`.)
* Methods with pointer receivers can **modify the value** to which the receiver points.
* Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
```

#### Pointers and functions

```go
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```
output: 50

```go
func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
```
output: 5

#### Methods and pointer indirection

* functions with a pointer argument must take a pointer
* methods with pointer receivers take either a value or a pointer as the receiver

```go
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)  // OK
	ScaleFunc(v, 10)  // Compile error
	ScaleFunc(&v, 10) // OK

	p := &Vertex{4, 3}
	p.Scale(3)  // OK
	ScaleFunc(p, 8)  // OK

	fmt.Println(v, p)
}
```

* Functions that take a value argument must take a value of that specific type

```go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())  // OK
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v))  // Compile error

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())  // (*p).Abs() OK
	fmt.Println(AbsFunc(*p))  // OK
	fmt.Println(AbsFunc(p))  // Compile error
}
```

#### Choosing a value or pointer receiver

* pointer receiver
    * so that the method can **modify the value** that its receiver points to.
    * to avoid copying the value on each method call. (**more efficient** if the receiver is a large struct ...)

* all methods on a given type should have either value or pointer receivers, but not a mixture of both.