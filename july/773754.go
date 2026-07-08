package july

func sumAndMultiply(n int) int64 {
	nums := make([]int, 0)
	for n > 0 {
		nums = append(nums, n%10)
		n = n / 10
	}
	j := 1
	sum := 0
	ne := 0
	for _, x := range nums {
		if x == 0 {
			continue
		}
		sum += x
		ne += x * j
		j *= 10
	}
	return int64(sum * ne)
}
