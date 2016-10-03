package pipelines

import (
	//"fmt"
	//"sync"
	"testing"
	//"time"
)

func TestPipelines(t *testing.T) {

	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	t.Log("Before calling Merge")

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		t.Log(n) // 4 then 9, or 9 then 4
	}

}
