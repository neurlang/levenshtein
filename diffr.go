package levenshtein

type pathStep struct {
	I, J int
	X, Y int
	T    byte
}

// levenpath takes a 1D slice representing the matrix and its width, and reconstructs the path
func levenpath[T Number](matrix []T, width uint) []pathStep {

	if len(matrix) == 0 {
		return nil
	}

	if *Distance(matrix) == 0 {
		return nil
	}

	// Reconstruct the optimal path from bottom-right to top-left
	var levenpath []pathStep
	height := uint(len(matrix) / int(width))
	j, i := int(width)-1, int(height)-1

	for {

		// Append the current step to the path
		levenpath = append(levenpath, pathStep{I: i, J: j})

		// Break if we've reached the top-left corner
		if i == 0 && j == 0 {
			break
		}

		// Handle the boundaries: when i == 0 or j == 0
		if i == 0 {
			j-- // Only move left if we're on the top row
		} else if j == 0 {
			i-- // Only move up if we're in the first column
		} else {

			var nextI, nextJ int

			var diag = matrix[int(width)*(i-1)+(j-1)]
			var up = matrix[int(width)*(i-1)+j]
			var left = matrix[int(width)*i+(j-1)]

			// Track back the path based on the minimum value (up, left, or diag)
			var min = diag
			if up < min {
				min = up
			}
			if left < min {
				min = left
			}
			// Update backtracking info based on which direction had the minimum cost
			if min == diag {
				nextI = i - 1
				nextJ = j - 1 // Came from diagonal (substitution)
			} else if left == min {
				nextI = i
				nextJ = j - 1 // Came from left (insertion)
			} else if up == min {
				nextI = i - 1
				nextJ = j // Came from above (deletion)
			}

			// If not at a boundary, move to the next backtracked step
			i, j = nextI, nextJ
		}
	}

	// Reverse the path to go from top-left to bottom-right
	for left, right := 0, len(levenpath)-1; left < right; left, right = left+1, right-1 {
		levenpath[left], levenpath[right] = levenpath[right], levenpath[left]
	}

	// Classify the operations (replace, insert, delete)
	for idx := 1; idx < len(levenpath); idx++ {

		prev := levenpath[idx-1]
		cur := levenpath[idx]

		// Access matrix values using 1D slice indexing
		prevVal := matrix[prev.I*int(width)+prev.J]
		curVal := matrix[cur.I*int(width)+cur.J]

		// Determine the type of edit
		if cur.I > prev.I && cur.J > prev.J && curVal != prevVal {

			levenpath[idx].X = cur.I - 1
			levenpath[idx].Y = cur.J - 1

			levenpath[idx].T = 3

		} else if cur.I == prev.I && cur.J > prev.J {

			levenpath[idx].X = cur.I
			levenpath[idx].Y = cur.J - 1

			levenpath[idx].T = 1

		} else if cur.I > prev.I && cur.J == prev.J {

			levenpath[idx].X = cur.I - 1
			levenpath[idx].Y = cur.J

			levenpath[idx].T = 2
		} else {
			levenpath[idx].X = cur.I
			levenpath[idx].Y = cur.J
		}

	}

	return levenpath
}

// Diff does diff based on edit distance matrix in reverse, detecting skips, inserts, deletes, replacements and their positions.
// Use false to stop the iteration.
// is_skip: True if the current characters from both sequences match.
// is_insert: True if a character from the second sequence is inserted.
// is_delete: True if a character from the first sequence is deleted.
// is_replace: True if a character from the first sequence is replaced with a character from the second sequence.
func Diff[T Number](mat []T, width uint, differ func(is_skip, is_insert, is_delete, is_replace bool, x, y uint) bool) {
	//height := uint(len(mat))/width
	var levenpath = levenpath(mat, width)

	for _, v := range levenpath {
		var is_skip bool
		var is_insert, is_delete, is_replace bool
		switch v.T {
		case 3:
			is_replace = true
		case 1:
			is_insert = true
		case 2:
			is_delete = true
		case 0:
			is_skip = true
		}
		if !differ(is_skip, is_insert, is_delete, is_replace, uint(v.X), uint(v.Y)) {
			return
		}
	}
}
