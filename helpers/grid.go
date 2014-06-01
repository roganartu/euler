package helpers

import (
	"math"
)

// HorizontalGridCombos returns a list of all horizontal combinations of n size
//
// The function assumes grid is square and that horizontal dimension is >= n.
// Only searches for combos left->right as this will find all.
func HorizontalGridCombos(grid [][]int, n int) [][]int {
	combos := make([][]int, 0)
	for _, row := range grid {
		sequence := make([]int, n)
		for j, cell := range row {
			sequence, _ = Shift(sequence)
			sequence[len(sequence)-1] = cell
			if j >= n-1 {
				combos = append(combos, sequence)
			}
		}
	}
	return combos
}

// VerticalGridCombos returns a list of all vertical combinations of n size
//
// The function assumes grid is square and that vertical dimension is >= n.
// Only searches for combos top->bottom as this will find all.
func VerticalGridCombos(grid [][]int, n int) [][]int {
	combos := make([][]int, 0)
	for i, row := range grid {
		if i+n >= len(grid) {
			break
		}

		for j, _ := range row {
			sequence := make([]int, n)
			for k := 0; k < n; k++ {
				sequence[k] = grid[i+k][j]
			}
			combos = append(combos, sequence)
		}
	}
	return combos
}

// DiagonalGridCombos returns a list of all diagonal combinations of n size
//
// The function assumes grid is square and that dimension is >= n.
// Only searches for left and right diagonals downwards as this will find all.
func DiagonalGridCombos(grid [][]int, n int) [][]int {
	combos := make([][]int, 0)
	for i, row := range grid {
		if i+n >= len(grid) {
			break
		}

		for j, _ := range row {
			if j+n >= len(grid) {
				break
			}
			for neg := -1; neg < 2; neg += 2 {
				if j-n < 0 {
					continue
				}

				sequence := make([]int, n)
				for k := 0; k < n; k++ {
					sequence[k] = grid[i+k][j+k*neg]
				}
				combos = append(combos, sequence)
			}
		}
	}
	return combos
}

// LatticePathCount returns the number of possible distinct paths through an nxn lattice.
//
// It utilises the fact that lattice path combinations are a binomial coefficient
// problem where the value of each point in the lattice is the sum of the node
// immediately above and left of it.
//
// O(n^2) time complexity.
// O(n^2) space complexity (could be simplified to O(n) with a circular buffer)
func LatticePathCount(n int) int {
	var j int
	bufSize := int(math.Pow(float64(n+1), 2))
	buffer := make([]int, bufSize)

	for i := 0; i < bufSize; i++ {
		j = i % bufSize
		if i%(n+1) == 0 || i < n+1 {
			buffer[j] = 1
		} else {
			buffer[j] = buffer[j-1] + buffer[j-n-1]
		}
	}
	return buffer[bufSize-1]
}
