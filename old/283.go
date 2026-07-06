package old

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if nums[i] == 0 {
			for j < len(nums) && nums[j] == 0 {
				j++
			}
			if j >= len(nums) {
				return
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
