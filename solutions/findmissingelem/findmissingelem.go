package findmissingelem

import "sort"

func Solution(a []int) int {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	if len(copyOfA) == 0 {
		return -1
	}

	if copyOfA[0] != 1 {
		return 1
	}

	for i := 0; i < len(copyOfA)-1; i++ {
		if copyOfA[i+1]-copyOfA[i] != 1 {
			return i + 2
		}
	}

	return -1
}
