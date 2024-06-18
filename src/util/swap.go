package util

func Swap(s []int, first, second int) {
	s[first], s[second] = s[second], s[first]
}
