package helpers

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestLatticePathCount(t *testing.T) {
	assert.Equal(t, 6, LatticePathCount(2), "Expected 2x2 grid to have 6 lattice paths")
	assert.Equal(t, 137846528820, LatticePathCount(20),
		"Expected 20x20 grid to have 137,846,528,820 lattice paths")
}
