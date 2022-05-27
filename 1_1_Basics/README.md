# Chapter 1. Basics
* [Packages](#packages)
    * [Imports](#imports)
    * [Exposed Name](#exposed-name)
* [Functions](#functions)
    * [Multiple results](#multiple-results)
    * [Named return values](#named-return-values)
* [Variables](#variables)
    * [Variables with initializers](#variables-with-initializers)
    * [Short variable declarations](#short-variable-declarations)
    * [Basic types](#basic-types)
    * [Zero values](#zero-values)
    * [Type conversions](#type-conversions)
    * [Type inference](#type-inference)
    * [Constants](#constants)
    * [Numeric Constants](#numeric-constants)

## Packages

* Every Go program is made up of packages.

* Programs start running in package main.

    ```go
    package main
    ```

### Imports

* Using the packages with import paths.
* "factored" import statement

    ```go
    import (
        "fmt"
        "math/rand"
    )
    ```

### Exposed Name

* A name is exported if it begins with a capital letter.
    * Any names which don't start with a capital letter are not exported.
    * Any unexported names are not accessible from outside the package.
* When importing a package, you can refer only to its exported names.

    ```go
    fmt.Println(math.Pi)
    fmt.Println("Time: ", time.Now())
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    fmt.Println("My favorite number is", rand.Intn(10))
    ```

## Functions

* The [type](https://www.geeksforgeeks.org/data-types-in-go/) comes after the variable name.
    * Numbers
        * Integers : `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`, \
            `int`, `uint` (Both int and uint contain same size, either 32 or 64 bit.), \
            `rune` (It is a synonym of int32 and also represent Unicode code points.), \
            `byte` (It is a synonym of uint8.), \
            `uintptr` (It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.)
        * Floating-Point Numbers : `float32`, `float64`
        * Complex Numbers : `complex32`, `complex64`
    * Booleans : `bool`
        ```go
        var result1 bool = true

        result1 := "111" == "111"
        ```
    * Strings : `string`
        ```go
        var result1 string = "string"

        result1 := "string"
        ```

```go
func add(x int, y int) int {
	return x + y
}

func get_string() string {
    return "golang"
}
```

* When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
```go
func add(x, y int) int {
	return x + y
}
```

* main function : `func main(int, []string) int`

### Multiple results

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

### Named return values

* These names should be used to document the meaning of the return values.
* Naked return statements should **be used only in short functions**, as with the example shown here. They can harm readability in longer functions.

```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
// >>> 7 10
```

## Variables

`var` statement

```go
var c, python, java bool
var i int
```

### Variables with initializers

```go
var c, python, java = true, false, "no!"
var i, j int = 1, 2
```

### Short variable declarations

The `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

```go
k := 3
c, python, java := true, false, "no!"
```

## Basic types

* Boolean: `bool`
* String: `string`
* Integer: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
* Byte: `byte`
* Rune: `rune`
* Float: `float32`, `float64`
* Complex: `complex64`, `complex128`

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

```go
import "math/cmplx"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1 << 64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
```

### Zero values

Variables declared without an explicit initial value are given their zero value.

* Number: `0`
* Boolean: `false`
* String: `""`

### Type conversions

The expression `T(v)` converts the value `v` to the type `T`.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

i := 42
f := float64(i)
u := uint(f)
```

### Type inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var =` expression syntax), the variable's type is inferred from the value on the right hand side.

```go
var a int
var b float32

c := 123 // int
d := 12.3 // float64
```

### Constants

`const`

> Constants cannot be declared using the `:=` syntax.

```go
const Pi = 3.14
```

### Numeric Constants

An untyped constant takes the type needed by its context.

```go
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
```