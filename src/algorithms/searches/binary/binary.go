package binary

func Search(s []int, elem int) int {
	return search(s, 0, len(s)-1, s[elem])
}

func search(s []int, start, end, elem int) int {
	if start <= end {
		mid := start + (end-start)/2

		if s[mid] == elem {
			return mid
		}

		if s[mid] > elem {
			return search(s, start, mid-1, elem)
		}

		return search(s, mid+1, end, elem)
	}

	return -1
}
