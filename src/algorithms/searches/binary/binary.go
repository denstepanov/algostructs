package binary

func Search(s []int, start, end, elem int) int {
	if start <= end {
		mid := start + (end-start)/2

		if s[mid] == elem {
			return mid
		}

		if s[mid] > elem {
			return Search(s, start, mid-1, elem)
		}

		return Search(s, mid+1, end, elem)
	}

	return -1
}
