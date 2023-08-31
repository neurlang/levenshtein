package levenshtein_test

import "fmt"
import "testing"
import "github.com/neurlang/levenshtein"

//Calculate the edit distance between unicode strings.

func TestUnicode(t *testing.T) {

	const word1 = "☺"
	const word2 = "Ö"

	var mat = levenshtein.Matrix[float32](uint(len([]rune(word1))), uint(len([]rune(word2))),
		nil, nil,
		levenshtein.OneSlice[rune, float32]([]rune(word1), []rune(word2)), nil)

	for x := 0; x <= len([]rune(word2)); x++ {
		for y := 0; y <= len([]rune(word1)); y++ {
			var pos = x + y*(len([]rune(word2))+1)
			fmt.Print(mat[pos], " ")
		}
		fmt.Println()
	}

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))

	if *levenshtein.Distance(mat) != 1 {
		t.Fail()
	}
}

//Calculate the edit distance between raw strings.

func TestRaw(t *testing.T) {

	const word1 = "☺"
	const word2 = "Ö"

	var mat = levenshtein.Matrix[float32](uint(len((word1))), uint(len((word2))),
		nil, nil,
		levenshtein.OneString[float32]((word1), (word2)), nil)

	for x := 0; x <= len((word2)); x++ {
		for y := 0; y <= len((word1)); y++ {
			var pos = x + y*(len((word2))+1)
			fmt.Print(mat[pos], " ")
		}
		fmt.Println()
	}

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))

	if *levenshtein.Distance(mat) != 3 {
		t.Fail()
	}
}

//Calculate the transposed edit distance between strings.

func TestTransposed(t *testing.T) {

	const word1 = "1234567"
	const word2 = "137"

	var mat = levenshtein.Matrix[float32](uint(len((word1))), uint(len((word2))),
		nil, nil,
		levenshtein.OneString[float32]((word1), (word2)), nil)
	var matt = levenshtein.MatrixT[float32](uint(len((word1))), uint(len((word2))),
		nil, nil,
		levenshtein.OneString[float32]((word1), (word2)), nil)

	for x := 0; x <= len((word2)); x++ {
		for y := 0; y <= len((word1)); y++ {
			var pos = x + y*(len((word2))+1)
			fmt.Print(mat[pos], " ")
		}
		fmt.Println()
	}
	for x := 0; x <= len((word1)); x++ {
		for y := 0; y <= len((word2)); y++ {
			var pos = x + y*(len((word1))+1)
			fmt.Print(matt[pos], " ")
		}
		fmt.Println()
	}

	for x := 0; x <= len((word2)); x++ {
		for y := 0; y <= len((word1)); y++ {
			var pos = x + y*(len((word2))+1)
			var post = y + x*(len((word1))+1)

			if mat[pos] != matt[post] {
				t.Fail()
			}
		}
	}

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))
	fmt.Println("Transposed Edit distance is:", *levenshtein.Distance(matt))

	if *levenshtein.Distance(mat) != *levenshtein.Distance(matt) {
		t.Fail()
	}
}

//Calculate the transposed edit distance between slices.

func TestTransposedSlices(t *testing.T) {

	var array1 = []string{"0", "1", "2"}
	var array2 = []string{"0", "2"}

	var matt = levenshtein.MatrixTSlices[float32, string](array1, array2,
		nil, nil, nil, nil)

	for x := 0; x <= len((array1)); x++ {
		for y := 0; y <= len((array2)); y++ {
			var pos = x + y*(len((array1))+1)
			fmt.Print(matt[pos], " ")
		}
		fmt.Println()
	}

	fmt.Println("Transposed Edit distance is:", *levenshtein.Distance(matt))
}
