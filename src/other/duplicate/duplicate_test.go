package duplicate

import "testing"

func TestHasArrayDuplicatesFalse(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := HasArrayDuplicates(arr)
	if result != false {
		t.Errorf("HasArrayDuplicates fail")
	}
}

func TestHasArrayDuplicatesTrue(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 10}
	result := HasArrayDuplicates(arr)
	if result != true {
		t.Errorf("HasArrayDuplicates fail")
	}
}
