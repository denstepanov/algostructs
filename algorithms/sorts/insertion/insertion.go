package insertion

func Sort(s []int) {
	len := len(s)

	maxIdx := len - 1
	for divider := 0; divider <= maxIdx; divider++ {
		for i := divider; i > 0; i-- {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}
}
