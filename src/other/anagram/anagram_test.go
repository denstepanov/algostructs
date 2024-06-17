package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	string1 := "kkaaz"
	string2 := "kazak"
	if !IsAnagram(string1, string2) {
		t.Errorf("anagram isn't anagram")
	}
}
