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
	
	for x := 0; x <= len([]rune(word2)); x++ {
	for y := 0; y <= len([]rune(word1)); y++ {
		var pos = x+y*(len([]rune(word2))+1)
		fmt.Print(mat[pos], " ")
	}
		fmt.Println()
	}

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
	
	for x := 0; x <= len((word2)); x++ {
	for y := 0; y <= len((word1)); y++ {
		var pos = x+y*(len((word2))+1)
		fmt.Print(mat[pos], " ")
	}
		fmt.Println()
	}

	fmt.Println("Edit distance is:", *levenshtein.Distance(mat))

}
```
