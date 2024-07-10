package main

import (
	"fmt"
	"sort"
)

func quicksort(list []int) []int {
	var sortedList []int

	if len(list) > 1 {
		var lessList []int
		var greaterList []int

		pivot := list[0]
		for i := 1; i < len(list); i++ {
			if list[i] < pivot {
				lessList = append(lessList, list[i])
			} else {
				greaterList = append(greaterList, list[i])
			}
		}

		fmt.Printf(" --> %v + %d + %v\n", lessList, pivot, greaterList)

		lessList = quicksort(lessList)
		greaterList = quicksort(greaterList)

		fmt.Printf(" --> %v + %d + %v\n", lessList, pivot, greaterList)

		sortedList = append(lessList, pivot)
		sortedList = append(sortedList, greaterList...)
	} else {
		sortedList = list
	}
	return sortedList
}

func main() {
	var listSize int
	fmt.Printf("Enter size of your list: ")
	fmt.Scanln(&listSize)

	var listNeedToSort = make([]int, listSize)

	for i := 0; i < listSize; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &listNeedToSort[i])
	}
	fmt.Printf("You want to sort %v.\n", listNeedToSort)

	sortedList := quicksort(listNeedToSort)

	// Checking the slice is sorted
	listIsSorted := sort.SliceIsSorted(sortedList, func(p, q int) bool {
		return sortedList[p] < sortedList[q]
	})

	if listIsSorted != true {
		fmt.Printf("%v is not sorted.\n", sortedList)
		return
	} else {
		fmt.Printf("%v is sorted\n", sortedList)
	}
}
