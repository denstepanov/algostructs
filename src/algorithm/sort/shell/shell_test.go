package shell_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithm/sort/shell"
)

const sortType = "Shell"

func TestDisorderedSlice(t *testing.T) {
	test.Sort(sortType, 10000, shell.Sort, t)
}
