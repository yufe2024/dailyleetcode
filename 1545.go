package main

import (
	"fmt"
	"strings"
)

func findKthBit(n int, k int) byte {
	var builder strings.Builder
	builder.Grow(2 ^ n)
	builder.WriteString("0")
	for i := 1; i < n; i++ {
		result := builder.String()
		builder.WriteString("1")
		builder.WriteString(re(reverse(result)))
	}
	result := builder.String()
	fmt.Println(result)
	return result[k]
}
func reverse(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(s); i++ {
		if bytes[i] == '0' {
			bytes[i] = '1'
		} else {
			bytes[i] = '0'
		}
	}
	return string(bytes)
}
func re(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
