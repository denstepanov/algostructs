package anagram

import "sort"

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	sortString(&str1)
	sortString(&str2)
	return str1 == str2
}

func sortString(str *string) {
	runes := []rune(*str)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	*str = string(runes)
}
