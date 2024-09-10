package levenshtein

// DiffR does diff based on edit distance matrix in reverse, detecting skips, inserts, deletes, replacements and their positions.
// Use false to stop the iteration.
func DiffR[T Number](mat []T, width uint, differ func(is_skip, is_insert, is_delete, is_replace bool, x, y uint) bool) {
	height := uint(len(mat)) / width
	var oldx, oldy = width - 1, height
	var distance = mat[len(mat)-1]
	WalkValsR(mat, width, func(prev, this T, x, y uint) bool {
		// Print matrix values
		if x == 0 || y == 0 {
			// break iteration when at the beginning of some word
			return true
		}

		if distance == this {

			var ret = differ(true, false, false, false, y, x)

			oldx = x
			oldy = y
			return ret
		}

		var is_insert, is_delete, is_replace bool
		switch [2]int{int(oldx) - int(x), int(oldy) - int(y)} {
		default:
			is_replace = true

		case [2]int{1, 0}, [2]int{-1, 0}:
			is_insert = true

		case [2]int{0, 1}, [2]int{0, -1}:
			is_delete = true

		}
		oldx = x
		oldy = y

		x--
		y--
		var cont = differ(false, is_insert, is_delete, is_replace, y, x)

		distance = this

		return cont

	})
}
