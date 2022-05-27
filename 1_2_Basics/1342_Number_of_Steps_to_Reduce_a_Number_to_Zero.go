// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/submissions/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
// Memory Usage: 1.9 MB, less than 83.90% of Go online submissions for Number of Steps to Reduce a Number to Zero.

package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	var step int = 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		step += 1
	}

	return step
}

func main() {
	fmt.Println(numberOfSteps(14)) // ans: 6
}
