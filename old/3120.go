//package main
//
//func numberOfSpecialChars(word string) int {
//	ans := 0
//	wd := []byte(word)
//	set := make(map[byte]bool)
//	for i := 0; i < len(wd); i++ {
//		set[wd[i]] = true
//	}
//	for c := byte('a'); c <= byte('z'); c++ {
//		if set[c] && set[c-32] {
//			ans++
//		}
//	}
//	return ans
//}
