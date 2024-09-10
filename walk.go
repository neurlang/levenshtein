package levenshtein

// Walk is the older method to iterate the edit distance matrix. It doesn't have any way to stop the iteration.
func Walk[T Number](mat []T, width uint, to func(uint, uint)) {
	WalkVals(mat, width, func(prev T, this T, x uint, y uint) bool {
		to(x, y)
		return false
	})
}

// WalkVals iterates the edit distance matrix. Use true to stop the iteration. X or Y may be zero even if one dimension was empty.
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
			if mat[diag] < mat[left] && mat[diag] < mat[up] {
				if to(mat[here], mat[diag], x, y) {
					return
				}
				x--
				y--
				continue
			}
		}
		if left >= 0 && up >= 0 {
			if mat[up] < mat[left] {
				if to(mat[here], mat[up], x, y) {
					return
				}
				y--
				continue
			} else {
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
	for x > 0 && y == 0 {
		here := x + width*y
		left := x - 1 + width*(y)
		if left >= 0 {
			if to(mat[here], mat[left], x, y) {
				return
			}
			x--
			continue
		}
		break
	}
	for x == 0 && y > 0 {
		here := x + width*y
		up := x + width*(y-1)
		if up >= 0 {
			if to(mat[here], mat[up], x, y) {
				return
			}
			y--
			continue
		}
		break
	}
}

// WalkVals iterates the edit distance matrix in reverse. Use false to stop the iteration.
// If any dimension is emtpy, the matrix is not walked.
func WalkValsR[T Number](mat []T, width uint, cb func(prev, this T, x, y uint) bool) {
	if width <= 1 {
		return
	}
	height := uint(len(mat)) / width
	if height <= 1 {
		return
	}
	WalkVals(mat, width, func(prev, this T, x, y uint) bool {
		x = width - x
		y = height - y
		return !cb(prev, this, x, y)
	})
}
