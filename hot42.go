package main

func trap(height []int) int {
	i := 0
	j := 1
	ans := 0
	n := len(height)
	for i < n-1 {
		for height[i] == 0 {
			i++
			j = i + 1
		}
		for height[j] == 0 {
			j++
		}
		if j == i+1 && height[j] > height[i] {
			i = j
			j = i + 1
			continue
		}
		if height[i] > height[j] {
			ans += height[j] * (j - i)
		}
	}
}
