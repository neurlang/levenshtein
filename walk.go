package levenshtein

// Walk is the older method to iterate the edit distance matrix. It doesn't have any way to stop the iteration. Using this for diff is broken, use Diff instead.
func Walk[T Number](mat []T, width uint, to func(uint, uint)) {
	WalkVals(mat, width, func(prev T, this T, x uint, y uint) bool {
		to(x, y)
		return false
	})
}

// WalkVals iterates the edit distance matrix. Use true to stop the iteration. Using this for diff is broken, use Diff instead.
func WalkVals[T Number](mat []T, width uint, to func(prev T, this T, x uint, y uint) bool) {
	pos := uint(len(mat) - 1)
	x := (pos % width)
	y := (pos / width)
	for x > 0 && y > 0 {
		here := x + width*y
		diag := x - 1 + width*(y-1)
		up := x + width*(y-1)
		left := x - 1 + width*(y)
		if diag >= 0 {

			if mat[left] == mat[up] && mat[diag] == mat[left] {
				if x == y {
					if to(mat[here], mat[diag], x, y) {
						return
					}
					x--
					y--
					continue
				}
				if x > y {
					if to(mat[here], mat[left], x, y) {
						return
					}
					x--
					continue
				}
				if x < y {
					if to(mat[here], mat[up], x, y) {
						return
					}
					y--
					continue
				}
			}

			if mat[diag] < mat[left] && mat[diag] < mat[up] {
				if to(mat[here], mat[diag], x, y) {
					return
				}
				x--
				y--
				continue
			}

			if mat[up] < mat[left] {
				if to(mat[here], mat[up], x, y) {
					return
				}
				y--
				continue
			}
			if mat[up] > mat[left] {
				if to(mat[here], mat[left], x, y) {
					return
				}
				x--
				continue
			}

		}
		if left >= 0 {
			if to(mat[here], mat[left], x, y) {
				return
			}
			x--
			continue
		}
		if up >= 0 {
			if to(mat[here], mat[up], x, y) {
				return
			}
			y--
			continue
		}
	}
}

// WalkVals iterates the edit distance matrix in reverse. Use false to stop the iteration. Using this for diff is broken, use Diff instead.
func WalkValsR[T Number](mat []T, width uint, cb func(prev, this T, x, y uint) bool) {
	height := uint(len(mat)) / width
	WalkVals(mat, width, func(prev, this T, x, y uint) bool {
		x = width - x
		y = height - y
		return !cb(prev, this, x, y)
	})
}
