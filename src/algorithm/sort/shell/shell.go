package shell

import "github.com/denstepanov/algostructs/src/util"

// Средний случай O(n log n)
// Худший случай O(n^2)
func Sort(s []int) {
	len := len(s)
	for step := len / 2; step > 0; step /= 2 {
		// Проход по элементам в рамках заданного шага.
		for i := step; i < len; i++ {
			// Сортировка элементов в рамках одной группы.
			// Групп может быть от двух и больше.
			// В конце сортировка переходит в сортировку вставками.
			for j := i; j >= step && s[j-step] > s[j]; j -= step {
				util.Swap(s, j-step, j)
			}
		}
	}
}
