// https://go.dev/tour/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	var t float64

	for { // while
		t = z
		z -= (z*z - x) / (2 * z)

		if math.Abs(t-z) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
