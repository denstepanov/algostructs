package quick

import "github.com/denstepanov/algostructs/src/utils"

// Сильно модифицированная версия пузырьковой сортировки.
// Лучший и средний случай O(n log n).
// Худший случай O(n^2). Массив частично упорядочен.
func Sort(s []int) {
	sort(s, 0, len(s)-1)
}

func sort(s []int, low, high int) {
	if low < high {
		pointer := partition(s, low, high)
		sort(s, low, pointer-1)
		sort(s, pointer+1, high)
	}
}

func partition(s []int, low, high int) int {
	pivot := high // индекс элемента-разделителя.

	pointer := low // позиция для элемента разделителя.
	for i := low; i < high; i++ {
		if s[i] < s[pivot] {
			s[pointer], s[i] = s[i], s[pointer]
			pointer++
		}
	}
	utils.Swap(s, pointer, pivot)
	return pointer
}
