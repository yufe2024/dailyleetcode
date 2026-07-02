package main

func maxArea(height []int) int {
	ans := 0
	i := 0
	j := len(height) - 1
	ans = min(height[i], height[j]) * (j - i)
	for i < j {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if min(height[i], height[j])*(j-i) > ans {
			ans = min(height[i], height[j]) * (j - i)
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
