package selection

// Во всех случаях O(n^2)
func Sort(s []int) {
	for start := 0; start < len(s)-1; start++ {
		minValueIndex := start
		for i := start; i < len(s); i++ {
			if s[i] < s[minValueIndex] {
				minValueIndex = i
			}
		}
		s[start], s[minValueIndex] = s[minValueIndex], s[start]
	}
}
