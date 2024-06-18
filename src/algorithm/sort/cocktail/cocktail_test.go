package cocktail_test

import (
	"github.com/denstepanov/algostructs/src/algorithm/sort/cocktail"
	"github.com/denstepanov/algostructs/src/test"
	"testing"
)

const sortType = "Cocktail"

func TestCocktailSortWith10kElements(t *testing.T) {
	test.Sort(sortType, 10000, cocktail.Sort, t)
}

func TestCocktailSortWithOneElement(t *testing.T) {
	test.Sort(sortType, 10000, cocktail.Sort, t)
}

func TestCocktailSortWithZeroElements(t *testing.T) {
	test.Sort(sortType, 10000, cocktail.Sort, t)
}
