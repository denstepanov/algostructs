package merge

// TODO: Подумать над in-place вариантом (без использования дополнительного массива)
/*
 * Быстрая сортировка слиянием
 *
 * Деление массива на две части. Затем деление частей на части до тех пор, пока изначальный массив не будет покрошен на отдельные элементы.
 * Затем эти элементы начинаем собирать.
 * Как только массивы будут собраны, следует проводить слияние следующих пар массивов, до тех пор пока не придём к итоговому массиву.
 *
 * Во всех случаях O(n log n)
 */
func Sort(s []int) {
	sorted := sort(s)
	// Замена элементов исходного массива на значения из отсортированного массива.
	for i := range s {
		s[i] = sorted[i]
	}
}

func sort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) / 2

	left := append([]int{}, s[:mid]...)
	right := append([]int{}, s[mid:]...)

	return merge(sort(left), sort(right))
}

func merge(left []int, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	totalLen := leftLen + rightLen
	leftPointer := 0
	rightPointer := 0

	result := make([]int, 0, totalLen)

	for i := 0; i < totalLen; i++ {
		if leftPointer < leftLen && rightPointer < rightLen {
			// Сравнение элементов из массивов и перемещение наименьшего значения в result.
			if left[leftPointer] > right[rightPointer] {
				result = append(result, right[rightPointer])
				rightPointer++
			} else {
				result = append(result, left[leftPointer])
				leftPointer++
			}
			// Закидывание оставшихся элементов из массивов.
		} else if rightPointer < rightLen {
			result = append(result, right[rightPointer])
			rightPointer++
		} else {
			result = append(result, left[leftPointer])
			leftPointer++
		}
	}
	return result
}
