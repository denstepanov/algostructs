package selection

import "github.com/denstepanov/algostructs/src/utils"

// Во всех случаях O(n^2)
func Sort(s []int) {
	for start := 0; start < len(s)-1; start++ {
		minValueIndex := start
		for i := start; i < len(s); i++ {
			if s[i] < s[minValueIndex] {
				minValueIndex = i
			}
		}
		utils.Swap(s, start, minValueIndex)
	}
}
