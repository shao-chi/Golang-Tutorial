# chapter 1. Basics

* [Pointer](#pointer)
* [Structs](#structs)
    * [Pointers to structs](#pointers-to-structs)
    * [Struct Literals](#struct-literals)
* [Arrays](#arrays)
* [Slices](#slices)
    * [Slices are like references to arrays](#slices-are-like-references-to-arrays)
    * [Slices Defaults](#slice-defaults)
    * [Nil slices](#nil-slices)
    * [Creating a slice with make](#creating-a-slice-with-make)
    * [Slices of slices](#slices-of-slices)
    * [Appending to a slice](#appending-to-a-slice)
    * [Range](#range)
* [Maps](#maps)
    * [Maps Literals](#map-literals)
    * [Mutating Maps](#mutating-maps)
* [Function values](#function-values)
* [Function closures](#function-closures)

> Exercises:
> 1. [Slices](/1_3_Basics/slices.go)
> 2. [Maps](/1_3_Basics/maps.go)
> 3. [Fibonacci closure](/1_3_Basics/fibonacci_closure.go)

### Pointer

A pointer holds the memory address of a value.
* The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
* The `&` operator generates a pointer to its operand.
* The `*` operator denotes the pointer's underlying value.
* Go has no pointer arithmetic.

```go
i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer
*p = 21         // set i through the pointer ("dereferencing" or "indirecting")
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j
```

### Structs

* A collection of fields.

```go
type Vertex struct {
	X int
	Y int
    // X, Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X)

	v.X = 4
	fmt.Println(v.X)
}
```

#### Pointers to structs

```go
func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println((*p).X)
	p.X = 1e9
	fmt.Println((*p).X)
	fmt.Println(v)
}
```
```plain test
1
1000000000
{1000000000 2}
```

#### Struct Literals

* A struct literal denotes a newly allocated struct value by listing the values of its fields.
* You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)
* The special prefix `&` returns a pointer to the struct value.

```go
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

### Arrays

* The type `[n]T` is an array of `n` values of type `T`.
* An array's length is part of its type, so arrays cannot be resized.

```go
var a [2]string
a[0] = "Hello"
a[1] = "World"

primes := [6]int{2, 3, 5, 7, 11, 13}
```

### Slices

* A slice is a dynamically-sized, flexible view into the elements of an array.
* The type `[]T` is a slice with elements of type `T`.
* A slice is formed by specifying two indices, a low and high bound, separated by a colon:
    ```go
    a[low : high]
    ```
* A slice literal is like an array literal without the length.
    This creates an array, then builds a slice that references it:
    ```go
    q := []int{2, 3, 5, 7, 11, 13}

	r := []bool{true, false, true, true, false, true}

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
    ```

```go
primes := [6]int{2, 3, 5, 7, 11, 13}  // array

var s []int = primes[1:4]  // slice
```

#### **Slices are like references to arrays**

* **A slice does not store any data, it just describes a section of an underlying array.**
* Changing the elements of a slice modifies the corresponding elements of its underlying array.
* Other slices that share the same underlying array will see those changes.

```go
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```
```plain text
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

#### Slice defaults

The default is zero for the low bound and the length of the slice for the high bound. \
    `a[:10]` is equivalent to `a[0:10]`

#### Slice length and capacity

* length : the number of elements it contains. `len(s)`
* capacity : the number of elements in the underlying array, counting from the first element in the slice. `cap(s)`
* You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

```go
s := []int{2, 3, 5, 7, 11, 13} // s = [2 3 5 7 11 13], len(s) == 6, cap(s) == 6

s = s[:0] // s = [], len(s) == 0, cap(s) == 6
s = s[:4] // s = [2 3 5 7], len(s) == 4, cap(s) == 6
s = s[2:] // s = [5 7], len(s) == 2, cap(s) == 4
```

#### Nil slices

* A nil slice has a length and capacity of 0 and has no underlying array.

```go
var s []int // s == nil, s: [], len(s) = 0, cap(s) = 0
fmt.Println(s == nil) // true
```

#### Creating a slice with make

* `make`: built-in function

```go
a := make([]int, 5)  // len(a) = 5, cap(a) = 5, a: [0 0 0 0 0]
b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5, b: []

c := b[:2] // len(c) = 2, cap(c) = 5, c: [0 0]
d := c[2:5] // len(d) = 3, cap(d) = 3, d: [0 0 0]
```

#### Slices of slices

* Slices can contain any type, including other slices.

```go
board := [][]string{
	[]string{"_", "_", "_"},  // [["_", "_", "_"],
	[]string{"_", "_", "_"},  //  ["_", "_", "_"],
	[]string{"_", "_", "_"},  //  ["_", "_", "_"]]
}
```

#### Appending to a slice

* `append`: built-in function `func append(s []T, vs ...T) []T`
* The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
* If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

```go
var s []int  // len(s) = 0, cap(s) = 0, s: []

s = append(s, 0)  // len(s) = 1, cap(s) = 1, s: [0]
s = append(s, 1)  // len(s) = 2, cap(s) = 2, s: [0, 1]
s = append(s, 2, 3, 4)  // len(s) = 5, cap(s) = 6, s: [0, 1, 2, 3, 4]
s = append(s, 1)  // len(s) = 6, cap(s) = 6, s: [0, 1, 2, 3, 4, 1]
s = append(s, 2, 3, 4)  // len(s) = 9, cap(s) = 12, s: [0, 1, 2, 3, 4, 1, 2, 3, 4]
```

#### Range

* The `range` form of the `for` loop iterates over a slice or map.
* When ranging over a slice, two values are returned for each iteration.
    1. index
    2. a copy of the element at that index

```go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

* skip the index or value by assigning to _
    ```go
    for i, _ := range pow
    for _, value := range pow
    ```

* only want the index
    ```go
    for i := range pow
    ```

### Maps

* A map maps keys to values.
* A `nil` map has no keys, nor can keys be added.
* The `make` function returns a map of the given type, initialized and ready for use.

```go
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

m = make(map[string]Vertex)
m["Bell Labs"] = Vertex{
	40.68433, -74.39967,
}
```

#### Map literals

* Map literals are like struct literals, but the keys are required.
* If the top-level type is just a type name, you can omit it from the elements of the literal.

```go
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
```

#### Mutating Maps

* Insert or update an element in map `m`: `m[key] = elem`
* Retrieve an element: `elem := m[key]`
* Delete an element: `delete(m, key)`
* Test that a key is present with a two-value assignment: `elem, ok := m[key]`
    * If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.
    * If `key` is not in the map, then `elem` is the zero value for the map's element type.

### Function values

* Functions are values too.
* Function values may be used as function arguments and return values.

```go
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

### Function closures

* A closure is a function value that references variables from outside its body.
* The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

```go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
```