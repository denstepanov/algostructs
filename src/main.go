package main

import "github.com/denstepanov/algostructs/src/other"

func main() {
	runOthers()
}

func runOthers() {
	other.RunReverseString("Hello World!")
	other.RunReverseSlice([]string{"Hello", ",", "World", "!"})
	other.RunIsPalindrome("kayak")
}
