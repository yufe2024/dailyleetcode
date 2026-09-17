package old

func rotate(nums []int, k int) {
	n := len(nums)
	num2 := make([]int, n)
	for i := 0; i < len(nums); i++ {
		num2[(i+k)%n] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = num2[i]
	}
}
