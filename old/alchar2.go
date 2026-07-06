package old

import (
	"fmt"
	"math"
)

func removeAlmostEqualCharacters(word string) int {
	count := 0
	i := 0
	for i < len(word)-1 {
		j := i
		char1 := int(word[j])
		char2 := int(word[j+1])
		fmt.Println(math.Abs(float64(char1 - char2)))
		if math.Abs(float64(char1-char2)) <= 1 {
			count++
			i += 2
		} else {
			i += 1
		}
	}
	return count
}
