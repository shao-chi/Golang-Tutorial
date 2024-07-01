package main

import (
	"fmt"
	"sort"
)

func findMinimumPosition(list []int) int {
	minimum := list[0]
	minimumPosition := 0
	for i := 1; i < len(list); i++ {
		if list[i] < minimum {
			minimum = list[i]
			minimumPosition = i
		}
	}
	return minimumPosition
}

func selectionSort(listNeedToSort []int) []int {
	sortedList := make([]int, len(listNeedToSort))
	listLength := len(listNeedToSort)
	for i := 0; i < listLength; i++ {
		minimumPosition := findMinimumPosition(listNeedToSort)
		sortedList[i] = listNeedToSort[minimumPosition]

		fmt.Printf(
			"- Step %d: %v found the minimum %d at position %d\n",
			i,
			listNeedToSort,
			sortedList[i],
			minimumPosition,
		)

		listNeedToSort = append(listNeedToSort[:minimumPosition], listNeedToSort[minimumPosition+1:]...)
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

	sortedList := selectionSort(listNeedToSort)

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
