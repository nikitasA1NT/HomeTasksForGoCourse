package cyclicrotation

import "sort"

//func Solution(a []int, k int) []int {
//
//}

func SingleShiftRight(a []int) []int {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	if len(copyOfA) != 0 {
		for i := len(copyOfA) - 1; i >= 1; i-- {
			copyOfA[i], copyOfA[i-1] = copyOfA[i-1], copyOfA[i]
		}
	}
	return copyOfA
}
