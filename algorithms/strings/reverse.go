package algorithms

func ReverseString(str string) string {
	if str == "" {
		return ""
	}

	runes := []rune(str)
	for head, tail := 0, len(runes)-1; head < tail; head, tail = head+1, tail-1 {
		runes[head], runes[tail] = runes[tail], runes[head]
	}
	return string(runes)
}

func IsPalindrome(str string) bool {
	if str == "" {
		return false
	}
	reversed := ReverseString(str)
	return str == reversed
}
