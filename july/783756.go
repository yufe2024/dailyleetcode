package july

func sumAndMultiply(s string, queries [][]int) []int {
	const mod int = 1e9 + 7
	ans := make([]int, 0)
	num := []byte(s)
	m := len(num)
	nums := make([]int, 0)
	pre := make([]int, m+1)
	sum := make([]int, m+1)
	cnt := make([]int, m+1)
	pow10 := make([]int, m+1)
	pow10[0] = 1
	nums = append(nums, 0)
	for i := 0; i < m; i++ {
		nums = append(nums, int(num[i])-'0')
	}
	for i := 1; i <= m; i++ {
		pow10[i] = pow10[i-1] * 10 % mod
		if nums[i] != 0 {
			pre[i] = (pre[i-1]*10 + nums[i]) % mod
			sum[i] = (sum[i-1] + nums[i]) % mod
			cnt[i] = cnt[i-1] + 1
		} else {
			pre[i] = pre[i-1]
			sum[i] = sum[i-1]
			cnt[i] = cnt[i-1]
		}
	}
	for i := 0; i < len(queries); i++ {
		li := queries[i][0]
		ri := queries[i][1]
		l := li + 1
		r := ri + 1
		totalSum := (sum[r] - sum[l-1] + mod) % mod
		k := cnt[r] - cnt[l-1]
		var x int
		if k > 0 {
			tmp := pre[l-1] * pow10[k] % mod
			x = (pre[r] - tmp + mod) % mod
		}
		res := x * totalSum % mod
		ans = append(ans, res)
	}
	return ans
}

//func sumAndMultiply(s string, queries [][]int) []int {
//	const mod int = 1e9 + 7
//	ans := make([]int, 0)
//	num := []byte(s)
//	nums := make([]int, 0)
//	pre := make([]int, len(num)+1)
//	sum := make([]int, len(num)+1)
//	nums = append(nums, 0)
//	for i := 0; i < len(num); i++ {
//		nums = append(nums, int(num[i])-'0')
//	}
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != 0 {
//			pre[i] = (pre[i-1]*10 + nums[i]) % mod
//			sum[i] = (sum[i-1] + nums[i]) % mod
//		} else {
//			pre[i] = pre[i-1]
//			sum[i] = sum[i-1]
//		}
//	}
//	for i := 0; i < len(queries); i++ {
//		l := queries[i][0]
//		r := queries[i][1]
//		ans = append(ans, (pre[r]-pre[l])*(sum[r]-sum[l])%mod)
//
//	}
//	return ans
//}
