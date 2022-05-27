package main

import (
	"fmt"
	"math"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

var (
	a int    = 2
	b int    = 1
	c bool   = true
	d string = "string"
)

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	sum := a + b // the variables with different types can't add
	e := "dddd"

	fmt.Println(sum)
	fmt.Println(c)
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Printf("d is of type %T\n", d)
	fmt.Println(e)

	var f float64 = math.Sqrt(float64(a*b + b*b))
	fmt.Println(f)
}
