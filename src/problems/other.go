package problems

import (
	"fmt"
	"github.com/denstepanov/algostructs/src/problems/reverse"
)

func RunIsPalindrome(word string) {
	fmt.Printf("\nIs this word '%s' a polindrome? %v", word, reverse.IsPalindrome(word))
}

func RunReverseString(word string) {
	fmt.Printf("\n Original string: %s", word)
	word = reverse.ReverseString(word)
	fmt.Printf("\n Reversed string: %s", word)
}

func RunReverseSlice(slice []string) {
	fmt.Printf("\nOriginal slice: %s", slice)
	slice = reverse.ReverseSliceOfStrings(slice)
	fmt.Printf("\nReversed slice: %s", slice)
}
