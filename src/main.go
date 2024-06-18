package main

import "github.com/denstepanov/algostructs/src/problems"

func main() {
	runOthers()
}

func runOthers() {
	problems.RunReverseString("Hello World!")
	problems.RunReverseSlice([]string{"Hello", ",", "World", "!"})
	problems.RunIsPalindrome("kayak")
}
