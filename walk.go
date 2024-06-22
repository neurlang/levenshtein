package levenshtein

func Walk[T Number](mat []T, width uint, to func(uint, uint)) {
	WalkVals(mat, width, func(prev T, this T, x uint, y uint) bool {
		to(x, y)
		return false
	})
}

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
}
