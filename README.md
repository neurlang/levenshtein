# levenshtein

Go implementation of the Levenshtein edit distance / diff algorithm with generics support.

```bash
go get github.com/neurlang/levenshtein
```

Requires Go 1.18 (generics) or later.

## What is Levenshtein Distance?

The Levenshtein distance measures the minimum number of single-element edits (insertions, deletions, or substitutions) needed to transform one sequence into another. It works with strings, slices, and any comparable types.

## Quick Start

### Calculate Edit Distance

```go
package main

import (
    "fmt"
    "github.com/neurlang/levenshtein"
)

func main() {
    word1 := []rune("kitten")
    word2 := []rune("sitting")
    
    mat := levenshtein.MatrixSlices[int](word1, word2, nil, nil, nil, nil)
    distance := levenshtein.Distance(mat)
    
    fmt.Println("Edit distance:", *distance) // Output: 3
}
```

### Diff Two Sequences

```go
func main() {
    word1 := []rune("demonstration")
    word2 := []rune("Demolition")
    
    mat := levenshtein.MatrixSlices[int](word1, word2, nil, nil, nil, nil)
    width := uint(len(word2) + 1)
    
    levenshtein.Diff(mat, width, func(is_skip, is_insert, is_delete, is_replace bool, x, y uint) bool {
        if is_skip {
            fmt.Printf("Skip: [%d][%d]\n", x, y)
        } else if is_replace {
            fmt.Printf("Replace: '%c' -> '%c' at [%d][%d]\n", word1[x], word2[y], x, y)
        } else if is_insert {
            fmt.Printf("Insert: '%c' at [%d][%d]\n", word2[y], x, y)
        } else if is_delete {
            fmt.Printf("Delete: '%c' at [%d][%d]\n", word1[x], x, y)
        }
        return true // continue iteration
    })
}
```

## API

### Core Functions

**`Matrix[T Number](m, n uint, deletion, insertion, substCost, kernel) []T`**  
Computes the Levenshtein distance matrix for two sequences of length m and n.

**`MatrixSlices[T Number, S comparable](a, b []S, deletion, insertion, substCost, kernel) []T`**  
Convenience wrapper for slice inputs.

**`MatrixT[T Number](...) []T`** / **`MatrixTSlices[T Number, S comparable](...) []T`**  
Transposed versions that swap source and target sequences.

**`Distance[T any](matrix []T) *T`**  
Returns the final edit distance from a computed matrix.

**`Diff[T Number](matrix []T, width uint, callback func(...) bool)`**  
Iterates through the edit path, calling the callback for each operation (skip, insert, delete, replace).

### Helper Functions

**`OneSlice[C comparable, T Number](a, b []C) func(uint, uint) *T`**  
Default substitution cost function for slices (returns 1 for different elements, nil for same).

**`OneString[T Number](a, b string) func(uint, uint) *T`**  
Default substitution cost function for raw strings (byte-level comparison).

**`OneElements[C comparable, T Number](a, b *C) *T`**  
Default substitution cost for individual elements.

## Examples

### Unicode Strings

```go
word1 := []rune("☺")
word2 := []rune("Ö")

mat := levenshtein.MatrixSlices[int](word1, word2, nil, nil, nil, nil)
fmt.Println("Distance:", *levenshtein.Distance(mat)) // Output: 1
```

### Raw Byte Strings

```go
word1 := "hello"
word2 := "hallo"

mat := levenshtein.Matrix[int](
    uint(len(word1)), uint(len(word2)),
    nil, nil,
    levenshtein.OneString[int](word1, word2),
    nil,
)
fmt.Println("Distance:", *levenshtein.Distance(mat)) // Output: 1
```

### Custom Types

```go
type1 := []string{"apple", "banana", "cherry"}
type2 := []string{"apple", "cherry"}

mat := levenshtein.MatrixSlices[int](type1, type2, nil, nil, nil, nil)
fmt.Println("Distance:", *levenshtein.Distance(mat)) // Output: 1
```

## Customization

All functions accept optional callbacks for:
- **deletion**: Custom deletion cost per element
- **insertion**: Custom insertion cost per element  
- **substCost**: Custom substitution cost between elements
- **kernel**: Custom algorithm kernel for cost calculation

Pass `nil` to use defaults (cost of 1 for all operations).

## License

See LICENSE file.
