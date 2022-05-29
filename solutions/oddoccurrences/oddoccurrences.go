package oddoccurrences

import (
	"errors"
	"sort"
)

func Solution(a []int) (result int, err error) {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	// If length of array is even
	if len(copyOfA)%2 == 0 {
		return 0, errors.New("empty array or its length is even")
	}

	resultWasFound := false
	for i := 0; i < len(copyOfA)-1; {
		// Check necessary condition, but result was already found
		if copyOfA[i] != copyOfA[i+1] && resultWasFound {
			return 0, errors.New("wrong array")
		}

		if copyOfA[i] != copyOfA[i+1] && !resultWasFound {
			resultWasFound = true
			result = copyOfA[i]
			i++ // Next elem step
		} else {
			i += 2 // Standard step
		}
	}

	if resultWasFound {
		return result, nil
	} else {
		// Return last elem because there are all pairs before it
		return copyOfA[len(copyOfA)-1], nil
	}
}
