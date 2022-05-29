package sequencecheck

import (
	"errors"
	"sort"
)

func Solution(a []int) (int, error) {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	// If array is empty
	if len(copyOfA) == 0 {
		return 0, errors.New("empty array")
	}

	counter := 1
	for i := 0; i < len(a); i++ {
		if copyOfA[i] != counter {
			return 0, nil // Array is not a sequence
		}
		counter++
	}

	// Array is a sequence
	return 1, nil
}
