package main

import (
	"fmt"
)

func factorial(num int) int {
	var result int
	if num == 1 {
		result = 1
		// fmt.Printf("1! = %d\n", result)
		fmt.Print("1")
	} else {
		result = num * factorial(num-1)
		// fmt.Printf("%d * %d! = %d\n", num, num-1, result)
		fmt.Printf(" * %d", num)
	}
	return result
}

func main() {
	var factorialNum int
	fmt.Printf("Enter factorial number: ")
	fmt.Scanln(&factorialNum)

	fmt.Printf("%d! = ", factorialNum)
	result := factorial(factorialNum)
	fmt.Printf(" = %d\n", result)
}
