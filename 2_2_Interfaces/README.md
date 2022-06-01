Chapter 2-2. Interfaces

* [Interfaces](#interfaces)
    * [Interfaces are implemented implicitly](#interfaces-are-implemented-implicitly)
    * [Interface values](#interface-values)
    * [Interface values with nil underlying values](#interface-values-with-nil-underlying-values)
    * [Nil interface values](#nil-interface-values)
    * [The empty interface](#the-empty-interface)
    * [Type assertions](#type-assertions)
    * [Type switches](#type-switches)
* [Stringers](#stringers)
* [Errors](#errors)
* [Readers](#readers)
* [Images](#images)

> Exercises:
> 1. [Stringers](/2_2_Interfaces/stringers.go)
> 2. [Errors](/2_2_Interfaces/errors.go)
> 3. [Readers](/2_2_Interfaces/readers.go)
> 4. [rot13Reader](/2_2_Interfaces/rot13Reader.go)
> 5. [Images](/2_2_Interfaces/images.go)

### Interfaces

* An interface type is defined as a set of method signatures.
* A value of interface type can hold any value that implements those methods.

```go
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

#### Interfaces are implemented implicitly

* A type implements an interface by implementing its methods.
* There is no explicit declaration of intent, no "implements" keyword.
* Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```go
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

#### Interface values

* An interface value holds a value of a specific underlying concrete type. `(value, type)`
* Calling a method on an interface value executes the method of the same name on its underlying type.

```go
type I interface {
	M()
}

type T struct {
	S string
}

// (value, type(struct))
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// (value, type(float64))
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}


// (value, type(interface))
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

#### Interface values with nil underlying values

* If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
* Note that an interface value that holds a nil concrete value is itself non-nil.

```go
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
```
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello
```

#### Nil interface values

* A nil interface value holds neither value nor concrete type.
* Calling a method on a nil interface is a **run-time error** because there is no type inside the interface tuple to indicate which concrete method to call.

```go
func main() {
	var i I
	describe(i)
	i.M()
}
```
`runtime error: invalid memory address or nil pointer dereference`

#### The empty interface

* the empty interface: The interface type that specifies zero methods `interface{}`
* An empty interface may hold values of any type. (Every type implements at least zero methods.)
* Empty interfaces are used by code that handles values of unknown type.
    * `fmt.Print` takes any number of arguments of type `interface{}`.

```go
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

#### Type assertions

* A type assertion provides access to an interface value's underlying concrete value.
* The interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

```go
// If `i` does not hold a `T`, the statement will trigger a panic.
t := i.(T)

// If `i` holds a `T`, then `t` will be the underlying value and `ok` will be true.
// If not, ok will be false and t will be the zero value of type T, and no panic occurs.
t, ok := i.(T)
```

```go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)  // 0, false

	f = i.(float64)  // panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)
}
```

#### Type switches

* A type switch is a construct that permits several type assertions in series.
* A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

```go
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

### Stringers

* One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package.
* A `Stringer` is a type that can describe itself as a string. **(Python: `__str__`)**
* The `fmt` package (and many others) look for this interface to print values.

```go
type Stringer interface {
    String() string
}
```

```go
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

### Errors

* The `error` type is a built-in interface similar to `fmt.Stringer`

```go
type error interface {
    Error() string  // the `fmt` package looks for the `error` interface when printing values.
}
```

* Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.
* A nil `error` denotes success; a non-nil `error` denotes failure.

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

```go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

### Readers

* The `io` package specifies the `io.Reader` interface, which represents the read end of a stream of data.
* The `io.Reader` interface has a `Read` method:

```go
func (T) Read(b []byte) (n int, err error)
```

* `Read` populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an `io.EOF` error when the stream ends.
* The Go standard library contains many implementations of this interface, including files, network connections, compressors, ciphers, and others.

```go
import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

### Images

* [Package image](https://go.dev/pkg/image/#Image) defines the `Image` interface:
	* Note: the `Rectangle` return value of the `Bounds` method is actually an `image.Rectangle`, as the declaration is inside package `image`.
	* The `color.Color` and `color.Model` types are also interfaces, but we'll ignore that by using the predefined implementations `color.RGBA` and `color.RGBAModel`. These interfaces and types are specified by the [image/color package](https://go.dev/pkg/image/color/)

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

```go
import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```