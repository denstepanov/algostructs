package anagram

import (
	"github.com/denstepanov/algostructs/src/util/strings"
)

func IsAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {

		return false
	}

	strings.Sort(&str1)
	strings.Sort(&str2)
	return str1 == str2
}
