package main

import (
	"fmt"

	algorithms "github.com/denstepanov/algostructs/algorithms/strings"
)

func main() {
	reverseString()
	reverseSlice()
	fmt.Printf("\nIs this word 'kazak' a polindrome? %v", algorithms.IsPalindrome("kazak"))
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
