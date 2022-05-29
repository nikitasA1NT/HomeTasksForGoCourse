package cyclicrotation

func Solution(a []int, k int) []int {
	copyOfA := make([]int, len(a))
	copy(copyOfA, a)

	for i := 0; i < k; i++ {
		SingleShiftRight(copyOfA)
	}

	return copyOfA
}

func SingleShiftRight(a []int) {
	for i := len(a) - 1; i >= 1; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}
