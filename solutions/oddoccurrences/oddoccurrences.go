package oddoccurrences

import "sort"

func Solution(a []int) int {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	if len(copyOfA)%2 == 0 {
		return -1
	}

	for i := 0; i < len(copyOfA)-1; i += 2 {
		if copyOfA[i] != copyOfA[i+1] {
			return copyOfA[i]
		}
	}
	return copyOfA[len(copyOfA)-1]
}
