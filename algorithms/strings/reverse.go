package algorithms

func ReverseString(str string) string {
	if str == "" {
		return ""
	}

	return string(reverse([]rune(str)))
}

func ReverseStringSlice(slc []string) []string {
	if len(slc) == 0 {
		return nil
	}
	return reverse(slc)
}

// Запилить более сложную версию определения палиндрома
func IsPalindrome(str string) bool {
	if str == "" {
		return false
	}
	reversed := ReverseString(str)
	return str == reversed
}

type runesOrStrings interface {
	rune | string
}

func reverse[T runesOrStrings](slice []T) []T {
	for head, tail := 0, len(slice)-1; head < tail; head, tail = head+1, tail-1 {
		slice[head], slice[tail] = slice[tail], slice[head]
	}
	return slice
}
