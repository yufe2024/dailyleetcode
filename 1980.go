package main

func findDifferentBinaryString(nums []string) string {
	var ans []byte
	lens := len(nums[0])
	for i := 0; i < lens; i++ {
		if nums[i][i] == '1' {
			ans = append(ans, '0')
		}
		if nums[i][i] == '0' {
			ans = append(ans, '1')
		}
	}
	return string(ans)
}
