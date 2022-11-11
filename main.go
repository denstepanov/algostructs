package main

import (
	"fmt"

	"github.com/denstepanov/algostructs/algorithms"
)

func main() {
	reverseString()
	reverseSlice()
	isPalindrome("kazak")
}

func reverseString() {
	word := "Hello, World!"
	fmt.Printf("\n Original string: %s", word)
	word = algorithms.ReverseString(word)
	fmt.Printf("\n Reversed string: %s", word)
}

func reverseSlice() {
	slice := []string{"Hello", ",", "World", "!"}
	fmt.Printf("\nOriginal slice: %s", slice)
	slice = algorithms.ReverseStringSlice(slice)
	fmt.Printf("\nReversed slice: %s", slice)
}

func isPalindrome(word string) {
	fmt.Printf("\nIs this word '%s' a polindrome? %v", word, algorithms.IsPalindrome("kazak"))
}
