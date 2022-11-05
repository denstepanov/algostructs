package algorithms

import "errors"

func ReverseString(str string) (string, error) {
	if str == "" {
		return "", errors.New("String is empty")
	}

	runes := []rune(str)
	for head, tail := 0, len(runes)-1; head < tail; head, tail = head+1, tail-1 {
		runes[head], runes[tail] = runes[tail], runes[head]
	}
	return string(runes), nil
}
