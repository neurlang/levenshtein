// Levenshtein implements the Levenshtein (edit distance) algorithm for golang
package levenshtein

// Number is the magnitude of the Levenshtein distance constraint
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
// or nil if the matrix is of size 0
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

// OneString instantiates the default substitution cost callback which calculates
// the substitution cost between two strings.
// The substitution cost is none if the string bytes are the same, and One if
// the string bytes are not equal.
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

// Kernel is the default Levenshtein algorithm kernel
func Kernel[T Number](d []T, i uint, j uint, n uint, cost *T) {
	var del = d[n*(i-1)+(j)] + d[n]
	var ins = d[n*(i)+(j-1)] + d[1]
	var sub = d[n*(i-1)+(j-1)] + *cost
	var cell = &d[n*(i)+(j)]

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

// Matrix is the default Levenshtein algorithm
func Matrix[T Number](m uint, n uint, deletion func(uint) *T, insertion func(uint) *T,
	substCost func(uint, uint) *T, kernel func(d []T, i uint, j uint, n uint, cost *T)) []T {

	m++
	n++

	var d = make([]T, m*n)

	var none = d[0]

	for i := uint(1); i < m; i++ {
		var dele *T
		if deletion != nil {
			dele = deletion(i - 1)
		}
		if dele == nil {
			d[n*(i)+0] = d[n*(i-1)+0] + none
		} else {
			d[n*(i)+0] = d[n*(i-1)+0] + *dele
		}
	}

	for j := uint(1); j < n; j++ {
		var inse *T
		if insertion != nil {
			inse = insertion(j - 1)
		}
		if inse == nil {
			d[n*(0)+j] = d[n*(0)+j-1] + none
		} else {
			d[n*(0)+j] = d[n*(0)+j-1] + *inse
		}
	}

	for j := uint(1); j < n; j++ {
		for i := uint(1); i < m; i++ {
			var cost *T
			if substCost != nil {
				cost = substCost(i-1, j-1)
			}
			if cost == nil {
				kernel(d, i, j, n, &none)
			} else {
				kernel(d, i, j, n, cost)
			}
		}
	}

	return d
}
