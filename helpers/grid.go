package helpers

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
