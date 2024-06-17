package quick

import "github.com/denstepanov/algostructs/src/utils"

// TODO: Придумать обёртку сортировки, чтобы была возможность переиспользовать общие методы для тестов
// Сильно модифицированная версия пузырьковой сортировки.
// Лучший и средний случай O(n log n).
// Худший случай O(n^2). Массив частично упорядочен.
func Sort(s []int, low, high int) {
	if low < high {
		pointer := partition(s, low, high)
		Sort(s, low, pointer-1)
		Sort(s, pointer+1, high)
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
