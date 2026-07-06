package main

import "fmt"

func minOperations(s string) int {
	count1 := 0
	count2 := 0
	count3 := 0
	count4 := 0
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] == '0' {
				count1++
			}
			if s[i] == '1' {
				count2++
			}
		}
		if i%2 == 1 {
			if s[i] == '0' {
				count3++
			}
			if s[i] == '1' {
				count4++
			}
		}
	}
	if count1+count4 < count3+count2 {
		return count1 + count4
	}
	fmt.Println(count2)
	fmt.Println(count3)
	return count3 + count2
}
func main() {
	var s string = "abddez"
	fmt.Println(removeAlmostEqualCharacters(s))
}
