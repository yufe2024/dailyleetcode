package october

func minSwaps(nums []int) int {
	num2 := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		an := 0
		for nums[i] > 0 {
			an = an * 10
			an += nums[i] % 10
			nums[i] = nums[i] / 10
		}
		num2 = append(num2, an)
	}

}
