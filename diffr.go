package levenshtein

// DiffR does diff based on edit distance matrix in reverse, detecting skips, inserts, deletes, replacements and their positions.
// Use false to stop the iteration.
// is_skip: True if the current characters from both sequences match.
// is_insert: True if a character from the second sequence is inserted.
// is_delete: True if a character from the first sequence is deleted.
// is_replace: True if a character from the first sequence is replaced with a character from the second sequence.
func DiffR[T Number](mat []T, width uint, differ func(is_skip, is_insert, is_delete, is_replace bool, x, y uint) bool) {
	height := uint(len(mat)) / width
	var oldx, oldy = width, height
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

		var is_skip, is_insert, is_delete, is_replace bool
		switch [2]bool{int(oldx)-int(x) != 0, int(oldy)-int(y) != 0} {
		case [2]bool{false, false}:
			is_insert = x > y
			is_delete = y > x
			is_skip = y == x
		case [2]bool{true, true}:
			is_insert = oldx > oldy
			is_delete = oldy > oldx
			is_replace = oldy == oldx

		case [2]bool{true, false}:
			is_insert = oldx > x
			is_replace = oldx < x

		case [2]bool{false, true}:
			is_delete = oldy > y
			is_replace = oldy < y
		}
		oldx = x
		oldy = y

		x--
		y--
		var cont = differ(is_skip, is_insert, is_delete, is_replace, y, x)

		distance = this

		return cont

	})
}
