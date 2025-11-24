package main

import "fmt"

func main() {
	unsorted := []int{5, 2, 4, 6, 1, 3}
	copyForIncreasing := append([]int(nil), unsorted...)
	copyForDecreasing := append([]int(nil), unsorted...)
	sortedIncreasing := insertionSort(copyForIncreasing, len(unsorted))
	sortedDecreasing := insertionSortDecreasing(copyForDecreasing, len(unsorted))
	fmt.Println(sortedIncreasing)
	fmt.Println(sortedDecreasing)
}

func insertionSort(unsortedArr []int, lenUnsortedArr int) []int {
	for i := 1; i < lenUnsortedArr; i++ {
		key := unsortedArr[i] // 2
		j := i - 1            // 0
		for j >= 0 && unsortedArr[j] > key {
			unsortedArr[j+1] = unsortedArr[j]
			j = j - 1
		}
		unsortedArr[j+1] = key
	}

	return unsortedArr
}

func insertionSortDecreasing(unsortedArr []int, lenUnsortedArr int) []int {
	for i := 1; i < lenUnsortedArr; i++ {
		key := unsortedArr[i] // 2
		j := i - 1            // 0
		for j >= 0 && unsortedArr[j] < key {
			unsortedArr[j+1] = unsortedArr[j]
			j = j - 1
		}
		unsortedArr[j+1] = key
	}

	return unsortedArr
}

// loop 1
/**
i = 1
key = 2
j = 0
unsortedArr[j] = 5 > 2
unsortedArr[1] = unsortedArr[0] = 5
j = -1
unsortedArr[0] = 2
**/

// loop 2
/**
key = 4
j = 1
unsortedArr[j] = 2 < 4
j = 0
unsortedArr[j] = 5 > 4
unsortedArr[1] = unsortedArr[0] = 5
j = -1
unsortedArr[j+1] = unsortedArr[-1+1] = unsortedArr[0] = 5
unsortedArr[0] = 4
unsortedArr[1] = 2
**/
