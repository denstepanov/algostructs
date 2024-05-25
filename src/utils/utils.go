package utils

import "fmt"

func IsPalindrome(word string) {
	fmt.Printf("\nIs this word '%s' a polindrome? %v", word, isPalindrome("kazak"))
}

func ReverseString() {
	word := "Hello, World!"
	fmt.Printf("\n Original string: %s", word)
	word = reverseString(word)
	fmt.Printf("\n Reversed string: %s", word)
}

func ReverseSlice() {
	slice := []string{"Hello", ",", "World", "!"}
	fmt.Printf("\nOriginal slice: %s", slice)
	slice = reverseStringSlice(slice)
	fmt.Printf("\nReversed slice: %s", slice)
}
