package bubble_test

import (
	"testing"

	"github.com/denstepanov/algostructs/algorithms"
	"github.com/denstepanov/algostructs/algorithms/sorts/bubble"
)

func TestEmptySlice(t *testing.T) {
	slice := []int{}

	_, tries := bubble.Sort(slice)

	if tries > 0 {
		t.Fatal("Bubble sort made an attempt to sort an empty slice.")
	}
}

func TestOneElemSlice(t *testing.T) {
	slice := []int{1}

	_, tries := bubble.Sort(slice)

	if tries > 0 {
		t.Fatal("Bubble sort made an attempt to sort an one elem slice.")
	}
}

func TestTwoElemSlice(t *testing.T) {
	slice := []int{2, 8}

	_, tries := bubble.Sort(slice)

	if tries != 1 {
		t.Fatal("Bubble sort didn't make an attempt to sort an two elem slice.")
	}

}

func TestOrderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(200)

	_, tries := bubble.Sort(ordered)

	if tries > 1 {
		t.Fatal("Bubble sort made an more than one attempt to sort an ordered slice.")
	}
}

func TestDisorderedSlice(t *testing.T) {
	ordered := algorithms.GenOrderedSlice(200)
	slice := algorithms.GenStirredSlice(ordered)

	slice, _ = bubble.Sort(slice)

	if !algorithms.SlicesAreEqual(ordered, slice) {
		t.Fatal("Bubble sort doesn't work correctly.")
	}
}
