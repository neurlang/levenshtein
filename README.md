# levenshtein and go slice diff algorithm
Levenshtein implements the Levenshtein (slice diff and edit distance) algorithm for golang (with generics) (go>=1.18)
```
import "github.com/neurlang/levenshtein"
```

## Overview

This package implements the Levenshtein (edit distance) algorithm using Go generics (requires Go >= 1.18).
The Levenshtein distance measures the minimum number of single-character edits (insertions, deletions, or
substitutions) required to change one word or sequence into another. This implementation extends the
algorithm to support slices and provides various utility functions to perform diffs, edit distance
calculations, and more.

## What is edit distance?

The Levenshtein distance is a measure of how dissimilar two sequences (such as strings or arrays) are by
counting the minimum number of edits needed to transform one sequence into the other. Edits can be insertions,
deletions, or substitutions. Skips occur when two characters do not have to be edited in order to transform
one sequence into the other (typically the same character in both sequences).

## Do diff between two slices (difference between two strings/slices in golang using generics)

Example:
```
Levenshtein edit distance between {1: 'demonstration.'} and {2: 'Demolition'} is:  7 edits

Levenshtein diff making {1: 'demonstration.'} into {2: 'Demolition'} (diff between two slices):

Edit at [ 0 ][ 0 ] 	swapped: {1: 'd'} in: {1: 'demonstration.'}: by {2: 'D'} of: {2: 'Demolition'}
Skip at [ 2 ][ 2 ] 
Skip at [ 3 ][ 3 ] 
Skip at [ 4 ][ 4 ] 
Edit at [ 4 ][ 3 ] 	deleted: {1: 'n'} in: {1: 'demonstration.'}: at {2: 'o'} of: {2: 'Demolition'}
Edit at [ 5 ][ 3 ] 	deleted: {1: 's'} in: {1: 'demonstration.'}: at {2: 'o'} of: {2: 'Demolition'}
Edit at [ 6 ][ 3 ] 	deleted: {1: 't'} in: {1: 'demonstration.'}: at {2: 'o'} of: {2: 'Demolition'}
Edit at [ 7 ][ 4 ] 	swapped: {1: 'r'} in: {1: 'demonstration.'}: by {2: 'l'} of: {2: 'Demolition'}
Edit at [ 8 ][ 5 ] 	swapped: {1: 'a'} in: {1: 'demonstration.'}: by {2: 'i'} of: {2: 'Demolition'}
Skip at [ 10 ][ 7 ] 
Skip at [ 11 ][ 8 ] 
Skip at [ 12 ][ 9 ] 
Skip at [ 13 ][ 10 ] 
Edit at [ 13 ][ 9 ] 	deleted: {1: '.'} in: {1: 'demonstration.'}: at {2: 'n'} of: {2: 'Demolition'}

```

In the following example, we use rune slices instead of raw strings to handle Unicode characters properly. This
is because Go's strings are byte-based, and using rune slices ensures that Unicode characters might be
interpreted as single elements.

Code:
```
func main() {
	// Define two words with some overlapping characters
	word1 := "demonstration."
	word2 := "Demolition"

	// Convert words to rune slices (to handle Unicode properly)
	w1r := []rune(word1)
	w2r := []rune(word2)

	// Create the Levenshtein edit distance matrix
	mat := levenshtein.Matrix[float32](
		uint(len(w1r)),
		uint(len(w2r)),
		nil, nil,
		levenshtein.OneSliceR[rune, float32](w1r, w2r),
		nil)

	var distance = *levenshtein.Distance(mat)

	// Now, print the Levenshtein edit distance
	println("\nLevenshtein edit distance between {1: '"+ word1+ "'} and {2: '"+ word2+ "'} is: ", int(distance), "edits")

	println("\nLevenshtein diff making {1: '"+ word1+ "'} into {2: '"+ word2+ "'} (diff between two slices):\n")

	// Finally do the diff of the two words
	levenshtein.DiffR(mat, uint(len(w2r)+1), func(is_skip, is_insert, is_delete, is_replace bool, x, y uint) bool {

		if is_skip {

			println("Skip at [", x, "][", y, "] ")

			return true
		}

		var operation, action string

		if is_replace {
			operation = "swapped"
			action = "by"
		}
		if is_insert {
			operation = "before"
			action = "put"
		}
		if is_delete {
			operation = "deleted"
			action = "at"
		}

		// Printing the position, words, and corresponding characters
		println("Edit at [", x, "][", y, "] \t" +
			operation+": {1: '"+ string(w1r[x])+ "'} in:",
				"{1: '"+ word1+ "'}"+":",
			action+" {2: '"+ string(w2r[y])+ "'} of:",
				"{2: '"+ word2+ "'}")

		// Continue walking through the matrix
		return true
	})

}
```

## Calculate the edit distance between unicode strings

```
func main1() {

	const word1 = "☺"
	const word2 = "Ö"

	var mat = levenshtein.Matrix[float32](uint(len([]rune(word1))), uint(len([]rune(word2))),
		nil, nil,
		levenshtein.OneSlice[rune, float32]([]rune(word1), []rune(word2)), nil)

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))

}
```

## Calculate the edit distance between raw strings

```
func main2() {

	const word1 = "☺"
	const word2 = "Ö"

	var mat = levenshtein.Matrix[float32](uint(len((word1))), uint(len((word2))),
		nil, nil,
		levenshtein.OneString[float32]((word1), (word2)), nil)

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))

}
```

## Calculate the transposed edit distance between string slices


```
func main3() {
	var array1 = []string{"0", "1", "2"}
	var array2 = []string{"0", "2"}

	var matt = levenshtein.MatrixTSlices[float32, string](array1, array2,
		nil, nil, nil, nil)

	fmt.Println("Transposed Edit distance is:", *levenshtein.Distance(matt))
}
```


