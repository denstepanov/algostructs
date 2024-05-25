package shell_test

import (
	"github.com/denstepanov/algostructs/src/test"
	"testing"

	"github.com/denstepanov/algostructs/src/algorithms/sorts/shell"
)

func TestDisorderedSlice(t *testing.T) {
	test.TestSort("Shell", shell.Sort, t)
}