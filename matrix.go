// Levenshtein implements the Levenshtein (edit distance and slice diff) algorithm for golang
package levenshtein

// Number is the magnitude of the Levenshtein distance constraint
// (the type such as integer which can be used to express the difference between characters or the two compared slices elements)
type Number interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | float32 | float64
}

// One returns the nonzero Levenshtein distance with magnitude of 1
func One[T Number](uint) *T {
	var n T
	n = 1
	return &n
}

// Distance returns the edit distance from the edit matrix (the last element),
// or nil if the matrix is of size 0. Used to figure out the Edit distance
// (how many minimum edits needed) from the matrix solved by Matrix.
func Distance[T any](d []T) *T {
	if len(d) == 0 {
		return nil
	}
	return &d[len(d)-1]
}

// OneSlice instantiates the default substitution cost callback which calculates
// the substitution cost between two comparable slices.
// The substitution cost is none if the slice elements are the same, and One if
// the slice elements are not equal.
// You can implement something like this to customize the substitution cost used by the slices algorithm.
func OneSlice[C comparable, T Number](a, b []C) func(uint, uint) *T {
	return func(x, y uint) *T {
		if a[x] == b[y] {
			return (*T)(nil)
		}
		var n T
		n = 1
		return &n
	}
}

// OneElements is the default substitution cost callback which calculates
// the substitution cost between two comparable elements.
// The substitution cost is none if the slice elements are the same, and One if
// the slice elements are not equal.
// You can implement something like this to customize the substitution cost used by the elements of slices algorithm.
func OneElements[C comparable, T Number](a *C, b *C) *T {
	if *a == *b {
		return (*T)(nil)
	}
	var n T
	n = 1
	return &n
}

// OneString instantiates the default substitution cost callback which calculates
// the substitution cost between two strings.
// The substitution cost is none if the string bytes are the same, and One if
// the string bytes are not equal.
// You can implement something like this to customize the substitution cost used by the strings algorithm.
func OneString[T Number](a, b string) func(uint, uint) *T {
	return func(x, y uint) *T {
		if a[x] == b[y] {
			return (*T)(nil)
		}
		var n T
		n = 1
		return &n
	}
}

// Kernel is the default Levenshtein algorithm kernel. You can customize the kernel to modify any costs or the algorithm calculation itself.
func Kernel[T Number](d []T, i uint, j uint, n uint, cost *T, delCost *T, insCost *T) {

	// Calculate deletion, insertion, and substitution costs
	var del = d[n*(i-1)+j] + *delCost
	var ins = d[n*i+(j-1)] + *insCost
	var sub = d[n*(i-1)+(j-1)] + *cost

	// Pointers to current cell and backtracking info
	var cell = &d[n*i+j]

	// Update cost in the current cell
	if del < ins {
		if del < sub {
			*cell = del
		} else {
			*cell = sub
		}
	} else {
		if ins < sub {
			*cell = ins
		} else {
			*cell = sub
		}
	}
}

// MatrixT is the transposed Levenshtein algorithm. Like Matrix except the the source sequence and the target sequence
// are exchanged. MatrixT returns a transposed matrix, while Matrix returns the original matrix.
func MatrixT[T Number](m uint, n uint, deletion func(uint) *T, insertion func(uint) *T,
	substCost func(uint, uint) *T, kernel func(d []T, i uint, j uint, n uint, cost *T, delCost *T, insCost *T)) []T {

	return Matrix(n, m, insertion, deletion, func(m uint, n uint) *T {
		return substCost(n, m)
	}, kernel)

}

// Matrix is the default Levenshtein algorithm. M is the size of the first sequence, N is the size of the second sequence.
// If the deletion, insertion, or substCost callbacks are not provided (i.e., passed as nil),
// the function uses default behaviors: insertions and deletions are treated as having a cost of 1,
// and substitutions are based on whether the characters are identical.
// The kernel actually returns the edit distance between two elements (you can customize this, or pass nil).
func Matrix[T Number](m uint, n uint, deletion func(uint) *T, insertion func(uint) *T,
	substCost func(uint, uint) *T, kernel func(d []T, i uint, j uint, n uint, cost *T, delCost *T, insCost *T)) []T {

	if kernel == nil {
		kernel = Kernel[T]
	}
	if deletion == nil {
		deletion = One[T]
	}
	if insertion == nil {
		insertion = One[T]
	}

	m++
	n++

	var d = make([]T, m*n)

	var none T

	for i := uint(1); i < m; i++ {
		var dele *T
		if deletion != nil {
			dele = deletion(i - 1)
		}
		if dele == nil {
			d[n*i+0] = d[n*(i-1)+0] + none
		} else {
			d[n*i+0] = d[n*(i-1)+0] + *dele
		}
	}

	for j := uint(1); j < n; j++ {
		var inse *T
		if insertion != nil {
			inse = insertion(j - 1)
		}
		if inse == nil {
			d[n*0+j] = d[n*0+j-1] + none
		} else {
			d[n*0+j] = d[n*0+j-1] + *inse
		}
	}

	for j := uint(1); j < n; j++ {
		for i := uint(1); i < m; i++ {
			var cost, delCost, insCost *T
			if substCost != nil {
				cost = substCost(i-1, j-1)
			}
			if deletion != nil {
				delCost = deletion(i - 1)
			}
			if insertion != nil {
				insCost = insertion(j - 1)
			}
			if cost == nil {
				cost = &none
			}
			if delCost == nil {
				delCost = &none
			}
			if insCost == nil {
				insCost = &none
			}
			kernel(d, i, j, n, cost, delCost, insCost)
		}
	}
	return d
}

// MatrixTSlices is the transposed Levenshtein algorithm for edit distance between slices. Use this instead of MatrixT,
// if your inputs are slices anyway. First two parameters are comparable slices, the other parameters can be nil or customized.
func MatrixTSlices[T Number, S comparable](ms, ns []S, deletion func(uint) *T, insertion func(uint) *T,
	substCost func(*S, *S) *T, kernel func(d []T, i uint, j uint, n uint, cost *T, delCost *T, insCost *T)) []T {

	if substCost == nil {
		substCost = OneElements[S, T]
	}

	return MatrixT(uint(len(ms)), uint(len(ns)), deletion, insertion, func(m uint, n uint) *T {
		return substCost(&ms[m], &ns[n])
	}, kernel)
}

// MatrixSlices is the Levenshtein algorithm for edit distance between slices. Use this instead of Matrix,
// if your inputs are slices anyway. First two parameters are comparable slices, the other parameters can be nil or customized.
func MatrixSlices[T Number, S comparable](ms, ns []S, deletion func(uint) *T, insertion func(uint) *T,
	substCost func(*S, *S) *T, kernel func(d []T, i uint, j uint, n uint, cost *T, delCost *T, insCost *T)) []T {

	if substCost == nil {
		substCost = OneElements[S, T]
	}

	return Matrix(uint(len(ms)), uint(len(ns)), deletion, insertion, func(m uint, n uint) *T {
		return substCost(&ms[m], &ns[n])
	}, kernel)
}
