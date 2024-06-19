package bubble

import "github.com/denstepanov/algostructs/src/util"

/*
 * Пузырьковая сортировка
 *
 * Проходя по массиву смотрим на два элемента. В самом начале это индексы 0 и 1.
 * Меняем местами, если левый элемент больше правого.
 * После первого прохода массива самый большой элемент будет располагаться в конце, что позволяет нам уменьшить длину прохода всего массива на 1.
 * Такое уменьшение массива похоже на оптимизацию. В оригинале алгоритм должен проходить массив от и до.
 *
* Лучший случай: O(n). Последовательность уже упорядочена.
 * Средний случай: O(n^2).
 * Худший случай: O(n^2). Последовательность обратно упорядочена.
 *
 * Показывает свою эффективность по сравнению с быстрой сортировкой если массив наполовину упорядочен.
*/
func Sort(s []int) {
	end := len(s) - 1
	var tries int
	for tries = 1; tries < len(s); tries++ {
		swaps := 0
		for left, right := 0, 1; right <= end; left, right = left+1, right+1 {
			if s[left] > s[right] {
				util.Swap(s, left, right)
				swaps++
			}
		}
		end--

		if swaps == 0 {
			break
		}
	}
}