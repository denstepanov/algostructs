package anagram

import (
	"github.com/denstepanov/algostructs/src/util/string"
)

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	string.Sort(&str1)
	string.Sort(&str2)
	return str1 == str2
}
