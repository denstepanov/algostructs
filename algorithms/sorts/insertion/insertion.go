package insertion

// Лучший случай O(n)
// Средний случай O(n^2)
// Худший случай O(n^2)
func Sort(s []int) {
	len := len(s)

	maxIdx := len - 1
	// Двигаю разделитель с начала и до конца, сортируя элементы левой части после каждого перемещения разделителя.
	for divider := 0; divider <= maxIdx; divider++ {
		for i := divider; i > 0; i-- {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}
}
