package oddoccurrences

import (
	"errors"
	"sort"
)

func Solution(a []int) (int, error) {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	// If length of array is even
	if len(copyOfA)%2 == 0 {
		return 0, errors.New("empty array or its length is even")
	}

	for i := 0; i < len(copyOfA)-1; i += 2 {
		if copyOfA[i] != copyOfA[i+1] {
			return copyOfA[i], nil
		}
	}
	// Return last elem because everything is pairs before it
	return copyOfA[len(copyOfA)-1], nil
}
