package cocktail

// Усовершенствованный вариант пузырьковой сортировки.
// Лучший вариант O(n). Срез отсортирован.
// Худший случай O(n^2). Срез обратно отсортирован.
func Sort(s []int) {
	if len(s) <= 1 {
		return
	}

	start := 0
	end := len(s) - 1
	for start < end {
		swaps := 0
		for i := start; i < end; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
				swaps++
			}
		}
		end--

		if swaps == 0 {
			break
		}

		for i := end; i > start; i-- {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
				swaps++
			}
		}
		start++

		if swaps == 0 {
			break
		}
	}
}
