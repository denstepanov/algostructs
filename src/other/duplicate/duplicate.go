package duplicate

func HasArrayDuplicates(s []int) bool {
	set := make(map[int]bool)
	for _, v := range s {
		set[v] = true
	}
	if len(set) < len(s) {
		return true
	}
	return false
}
