# levenshtein
Levenshtein implements the Levenshtein (edit distance) algorithm for golang (with generics) (go>=1.18)
```
import "github.com/neurlang/levenshtein"
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

## Do diff between two slices (difference between two strings in golang using generics)

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


