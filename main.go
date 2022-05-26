package main

import (
	"fmt"
	"sort"
)

func FindMissingElem(nums []int) (searchElem int) {
	copyOfNums := make([]int, len(nums))
	copy(copyOfNums, nums)

	sort.Ints(copyOfNums)

	if len(copyOfNums) == 0 {
		return -1
	}

	if copyOfNums[0] != 1 {
		return 1
	}

	for i := 0; i < len(copyOfNums)-1; i++ {
		if copyOfNums[i+1]-copyOfNums[i] != 1 {
			return i + 2
		}
	}

	return -1
}

func main() {
	// Tests of find a missing element in array
	slice1 := []int{2, 3, 1, 5}
	fmt.Println(FindMissingElem(slice1))

	slice2 := []int{7, 2, 3, 1, 5, 4}
	fmt.Println(FindMissingElem(slice2))

	slice3 := []int{2, 3, 4, 5, 6}
	fmt.Println(FindMissingElem(slice3))

	slice4 := []int{1, 2, 3, 4, 6}
	fmt.Println(FindMissingElem(slice4))

	slice5 := []int{}
	fmt.Println(FindMissingElem(slice5))
}
