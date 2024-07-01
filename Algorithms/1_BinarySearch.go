package main

import (
	"fmt"
	"sort"
)

func binarySearch(sortedList []int, targetNum int) int {
	fmt.Printf("Sorted List: %v\n", sortedList)
	fmt.Printf("Target number: %d\n", targetNum)

	start := 0
	end := len(sortedList) - 1
	mid := (start + end) / 2

	for start <= end {
		fmt.Printf("%v ... Found middle number: %d\n", sortedList[start:end+1], sortedList[mid])
		switch {
		case sortedList[mid] == targetNum:
			return mid
		case sortedList[mid] < targetNum:
			start = mid + 1
		default:
			// sortedList[mid] > targetNum
			end = mid - 1
		}

		mid = (start + end) / 2
	}
	return -1
}

func main() {
	var sortedListSize int
	fmt.Printf("Enter size of your sorted list: ")
	fmt.Scanln(&sortedListSize)

	var sortedList = make([]int, sortedListSize)

	for i := 0; i < sortedListSize; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &sortedList[i])
	}

	// Checking the slice is sorted
	listIsSorted := sort.SliceIsSorted(sortedList, func(p, q int) bool {
		return sortedList[p] < sortedList[q]
	})

	if listIsSorted != true {
		fmt.Println("Slice is not sorted.")
		return
	}

	var targetNum int
	fmt.Printf("Enter the target number you want to find: ")
	fmt.Scanln(&targetNum)

	position := binarySearch(sortedList, targetNum)
	if position == -1 {
		fmt.Printf("%d is not in %v\n", targetNum, sortedList)
	} else {
		fmt.Printf("The position in %v is %d\n", sortedList, position)
	}
}
