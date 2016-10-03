package pipelines

import (
	//"fmt"
	"testing"
)

func TestPipelines(t *testing.T) {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	t.Log(<-out) // 4
	t.Log(<-out) // 9
}

func TestPipelines2(t *testing.T) {
	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3))) {
		t.Log(n) // 16 then 81
	}
}

func TestPipelines3(t *testing.T) {
	// Set up the pipeline and consume the output.
	for n := range sq(gen(2, 3)) {
		t.Log(n) // 4 then 9
	}
}
