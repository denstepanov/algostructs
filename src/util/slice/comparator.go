package slice

func AreEqual[T comparable](a, b []T) bool {
	result := true
	for i, v := range a {
		if v != b[i] {
			result = false
		}
	}
	return result
}
