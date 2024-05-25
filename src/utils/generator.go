package utils

import (
	"math/rand"
	"time"
)

func GenStirredSlice(s []int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, len(s))
	copy(result, s)
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	return result
}

func GenOrderedSlice(size int) []int {
	ints := make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = i
	}
	return ints
}
