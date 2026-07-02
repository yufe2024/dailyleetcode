package main

func checkOnesSegment(s string) bool {
	var i int = 0
	for ; i < len(s); i++ {
		if s[i] == '0' {
			break
		}
	}
	for ; i < len(s); i++ {
		if s[i] == '1' {
			return false
		}
	}
	return true
}
