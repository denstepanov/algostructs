package main

import (
	"github.com/denstepanov/algostructs/src/utils"
)

func main() {
	utils.RunReverseString("Hello World!")
	utils.RunReverseSlice([]string{"Hello", ",", "World", "!"})
	utils.RunIsPalindrome("kayak")
}
