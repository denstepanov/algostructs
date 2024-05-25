package bubble

/*
 * Пузырьковая сортировка.
 * Лучший случай: O(n). Последовательность уже упорядочена.
 * Средний случай: O(n^2).
 * Худший случай: O(n^2). Последовательность обратно упорядочена.
 */
func Sort(s []int) {
	end := len(s) - 1
	var tries int
	for tries = 1; tries < len(s); tries++ {
		swaps := 0
		for left, right := 0, 1; right <= end; left, right = left+1, right+1 {
			if s[left] > s[right] {
				s[left], s[right] = s[right], s[left]
				swaps++
			}
		}
		end--

		if swaps == 0 {
			break
		}
	}
}
