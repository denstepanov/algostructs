package quick

// не подходит для частично упорядоченного массива. В этом случае O(n^2)
// по своей сути является модифицированной сортировкой пузырьком
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
	s[pointer], s[pivot] = s[pivot], s[pointer]
	return pointer
}
