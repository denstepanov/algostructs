package insertion

import "github.com/denstepanov/algostructs/src/utils"

// Лучший случай O(n)
// Средний случай O(n^2)
// Худший случай O(n^2)
func Sort(s []int) {
	if len(s) < 2 {
		return
	}

	lastIndex := len(s) - 1
	// Двигаю разделитель с начала и до конца, сортируя элементы левой части после каждого перемещения разделителя.
	for divider := 0; divider <= lastIndex; divider++ {
		// Цикл отвечает что каждое новое значение левого подмассива будет находиться в своём месте.
		// Для этого в качестве стартовой позиции указываем current == divider, затем уменьшаем значение current, пока оно не станет равно нулю.
		// Т.е. пробегаем левый массив в обратном порядке, меняя элементы местами, если предыдущий элемент больше текущего.
		for current := divider; current > 0; current-- {
			if s[current] < s[current-1] {
				utils.Swap(s, current, current-1)
			}
		}
	}
}
