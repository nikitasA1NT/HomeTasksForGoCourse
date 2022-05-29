package findmissingelem

import (
	"errors"
	"sort"
)

func Solution(a []int) (result int, err error) {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	sort.Ints(copyOfA)

	// If array is empty
	if len(copyOfA) == 0 {
		return 0, errors.New("empty array")
	}

	counter := 1
	resultWasFound := false
	for i := 0; i < len(copyOfA); i++ {
		// Check necessary condition, but result was already found
		if copyOfA[i] != counter && resultWasFound {
			return 0, errors.New("wrong array")
		}

		if copyOfA[i] != counter && !resultWasFound {
			resultWasFound = true
			result = counter // Current number
			counter += 2     //
		} else {
			counter++ // Standard step
		}
	}

	if resultWasFound {
		return result, nil
	} else {
		return 0, errors.New("wrong array")
	}
}
