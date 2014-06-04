package problems

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestProblem_3(t *testing.T) {
	assert.Equal(t, uint64(6857), Problem_3(uint64(600851475143)))
}

func TestProblem_5(t *testing.T) {
	assert.Equal(t, uint64(232792560), Problem_5())
}

func TestProblem_6(t *testing.T) {
	assert.Equal(t, uint64(25164150), Problem_6())
}

func TestProblem_7(t *testing.T) {
	assert.Equal(t, uint64(104743), Problem_7())
}

func TestProblem_8(t *testing.T) {
	assert.Equal(t, uint64(23514624000), Problem_8())
}

func TestProblem_9(t *testing.T) {
	assert.Equal(t, 31875000, Problem_9())
}

func TestProblem_10(t *testing.T) {
	assert.Equal(t, uint64(142913828922), Problem_10())
}

func TestProblem_11(t *testing.T) {
	assert.Equal(t, 70600674, Problem_11())
}

func TestProblem_17(t *testing.T) {
	assert.Equal(t, 21124, Problem_17())
}
