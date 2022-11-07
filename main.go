package main

import (
	"fmt"

	algorithms "github.com/denstepanov/algostructs/algorithms/strings"
)

func main() {
	execReverseString()
	fmt.Printf("\nIs this word 'kazak' a polindrome? %v", algorithms.IsPalindrome("kazak"))
}

func execReverseString() {
	word := "Hello, World!"
	fmt.Printf("\n Original string: %s", word)
	word = algorithms.ReverseString(word)
	fmt.Printf("\n Reversed string: %s", word)
}
