package october

func minOperations(s string) int {
	n := len(s)
	chars := []byte(s)
	nums := make([]int, len(chars))
	for i := 0; i < len(chars); i++ {
		nums[i] = int(chars[i] - '0')
	}
	ma := nums[0]
	countma := 0
	countmi := 0
	mi := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == mi {
			countmi++
		}
		if nums[i] < mi {
			mi = nums[i]
			countmi = 1
		}
		if nums[i] == ma {
			countma++
		}
		if nums[i] > ma {
			ma = nums[i]
			countma = 1
		}
	}
	right := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= nums[i-1] {
			continue
		} else {
			right = 0
		}
	}
	if right == 1 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return -1
		} else {
			return 0
		}
	}
	if nums[0] == mi || nums[n-1] == ma {
		return 1
	}
	if nums[0] == ma && nums[n-1] == mi {
		if countmi > 1 || countma > 1 {
			return 2
		}
		return 3
	}
	return 2
}
