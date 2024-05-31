package utils

import (
	"math/rand"
	"time"
)

func GenStirredSlice(s []int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, len(s))
	copy(list, s)
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
	return list
}

func GenOrderedSlice(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}
	return list
}
