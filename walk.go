package levenshtein

func Walk[T Number](mat []T, width uint, to func(uint, uint)) {
	pos := uint(len(mat) - 1)
	x := (pos % width)
	y := (pos / width)
	for x > 0 && y > 0 {
		diag := x - 1 + width*(y-1)
		up := x + width*(y-1)
		left := x - 1 + width*(y)
		if diag >= 0 {
			if mat[diag] < mat[left] && mat[diag] < mat[up] {
				x--
				y--
				to(x, y)
				continue
			}
		}
		if left >= 0 && up >= 0 {
			if mat[up] < mat[left] {
				y--
				to(x, y)
				continue
			} else {
				x--
				to(x, y)
				continue
			}
		}
		if left >= 0 {
			x--
			to(x, y)
			continue
		}
		if up >= 0 {
			y--
			to(x, y)
			continue
		}
	}
}
