package merge

func Sort(s []int) {
	sorted := sort(s)
	for i := range s {
		s[i] = sorted[i]
	}
}

func sort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2

	left := append([]int{}, s[:mid]...)
	right := append([]int{}, s[mid:]...)

	return merge(sort(left), sort(right))
}

func merge(left []int, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	totalLen := leftLen + rightLen
	leftPointer := 0
	rightPointer := 0

	result := make([]int, 0, totalLen)

	for i := 0; i < totalLen; i++ {
		if leftPointer < leftLen && rightPointer < rightLen {
			if left[leftPointer] > right[rightPointer] {
				result = append(result, right[rightPointer])
				rightPointer++
			} else {
				result = append(result, left[leftPointer])
				leftPointer++
			}
		} else if rightPointer < rightLen {
			result = append(result, right[rightPointer])
			rightPointer++
		} else {
			result = append(result, left[leftPointer])
			leftPointer++
		}
	}
	return result
}
