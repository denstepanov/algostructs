package reverse

import (
	"strings"
)

type runeOrString interface {
	rune | string
}

func IsPalindrome(str string) bool {
	if str == "" {
		return true
	}
	reversed := ReverseString(str)
	return str == reversed
}

func IsSliceOfStringPalindrome(slice []string) bool {
	if len(slice) == 0 {
		return true
	}
	reversedSlice := reverse(slice)
	sliceAsString := strings.Join(slice, "")
	reversedSliceAsString := strings.Join(reversedSlice, "")
	return sliceAsString == reversedSliceAsString
}

func ReverseString(str string) string {
	if len(str) == 0 {
		return str
	}
	return string(reverse([]rune(str)))
}

func ReverseSliceOfStrings(slc []string) []string {
	if len(slc) == 0 {
		return slc
	}
	return reverse(slc)
}

func reverse[T runeOrString](slice []T) []T {
	for head, tail := 0, len(slice)-1; head < tail; head, tail = head+1, tail-1 {
		slice[head], slice[tail] = slice[tail], slice[head]
	}
	return slice
}
