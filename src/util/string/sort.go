package string

import "sort"

func Sort(s *string) {
	runes := []rune(*s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	*s = string(runes)
}
