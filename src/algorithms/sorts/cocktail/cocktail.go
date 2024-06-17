package cocktail

import "github.com/denstepanov/algostructs/src/utils"

// Усовершенствованный вариант пузырьковой сортировки.
// Лучший случай O(n). Срез отсортирован.
// Средний случай O(n^2).
// Худший случай O(n^2). Срез обратно отсортирован.
func Sort(s []int) {
	start := 0
	end := len(s) - 1
	for start < end {
		swaps := 0
		for i := start; i < end; i++ {
			if s[i] > s[i+1] {
				utils.Swap(s, i, i+1)
				swaps++
			}
		}
		end--

		if swaps == 0 {
			break
		}

		for i := end; i > start; i-- {
			if s[i] < s[i-1] {
				utils.Swap(s, i, i-1)
				swaps++
			}
		}
		start++

		if swaps == 0 {
			break
		}
	}
}
