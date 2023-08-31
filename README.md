# levenshtein
Levenshtein implements the Levenshtein (edit distance) algorithm for golang (with generics) (go>=1.18)
```
import "github.com/neurlang/levenshtein"
```

Calculate the edit distance between unicode strings.

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

Calculate the edit distance between raw strings.

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

Calculate the transposed edit distance between string slices.

```
func main3() {
	var array1 = []string{"0", "1", "2"}
	var array2 = []string{"0", "2"}
	
	var matt = levenshtein.MatrixTSlices[float32, string](array1, array2,
		nil, nil, nil, nil)

	fmt.Println("Transposed Edit distance is:", *levenshtein.Distance(matt))
}
```
