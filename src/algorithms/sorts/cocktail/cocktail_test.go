package cocktail_test

import (
	"github.com/denstepanov/algostructs/src/algorithms/sorts/cocktail"
	"github.com/denstepanov/algostructs/src/test"
	"testing"
)

func TestCocktailSort(t *testing.T) {
	test.TestSort("Cocktail", cocktail.Sort, t)
}
