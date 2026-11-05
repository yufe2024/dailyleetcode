package novemember

import "sort"

func minOperations(nums []int, x int, k int) int64 {
	//中位数操作次数最少
	n := len(nums)
	an := make([]int, 0)
	dp := make([]int, n)
	for i := 0; i < n-x; i++ {
		linshi := make([]int, n-x)
		for j := 0; j < x; j++ {
			linshi[j] = nums[j+i]
		}
		sort.Ints(linshi)
		mid := linshi[len(linshi)/2]
		aa := 0
		for i := 0; i < len(linshi); i++ {
			aa += abs(mid - linshi[i])
		}
		an = append(an, aa)
	}
	//假设有了一串，看怎么不重叠求最小和,dp可以吗，选了之后后面x个不能选，看最后能不能够x
	//状态转移怎么做`
	for i := n - x; i < n; i++ {
		dp[i] = 0
	}
	for i := n - 1; i >= 0; i-- {
		dp[i] = min(dp[i+k], dp[i])
	}
	return int64(dp[n-x])
}

//import "sort"

//// 树状数组
//type fenwick struct {
//	tree []int64
//}
//
//func newFenwick(n int) *fenwick {
//	return &fenwick{tree: make([]int64, n+1)}
//}
//
//func (f *fenwick) add(i int, delta int64) {
//	for i < len(f.tree) {
//		f.tree[i] += delta
//		i += i & -i
//	}
//}
//
//// 求前 i 个位置之和
//func (f *fenwick) prefix(i int) int64 {
//	var sum int64
//	for i > 0 {
//		sum += f.tree[i]
//		i -= i & -i
//	}
//	return sum
//}
//
//// 仅用于计数树：找到第 k 小元素所在的位置
//func (f *fenwick) kth(k int64) int {
//	pos := 0
//	step := 1
//	for step<<1 < len(f.tree) {
//		step <<= 1
//	}
//
//	for ; step > 0; step >>= 1 {
//		next := pos + step
//		if next < len(f.tree) && f.tree[next] < k {
//			pos = next
//			k -= f.tree[next]
//		}
//	}
//	return pos + 1
//}

//func minOperations(nums []int, x int, k int) int64 {
//	n := len(nums)
//
//	// 离散化：将数值映射到从 1 开始的下标
//	values := append([]int(nil), nums...)
//	sort.Ints(values)
//
//	unique := values[:0]
//	for _, v := range values {
//		if len(unique) == 0 || unique[len(unique)-1] != v {
//			unique = append(unique, v)
//		}
//	}
//
//	rank := make(map[int]int, len(unique))
//	for i, v := range unique {
//		rank[v] = i + 1
//	}
//
//	// 按题目要求，在函数中途保存输入
//	maritovexi := struct {
//		nums []int
//		x, k int
//	}{nums, x, k}
//
//	countTree := newFenwick(len(unique))
//	sumTree := newFenwick(len(unique))
//	var totalSum int64
//
//	// 加入或删除一个元素，delta 分别为 1 或 -1
//	update := func(v int, delta int64) {
//		p := rank[v]
//		countTree.add(p, delta)
//		sumTree.add(p, int64(v)*delta)
//		totalSum += int64(v) * delta
//	}
//
//	cost := make([]int64, n-x+1)
//
//	for i, v := range maritovexi.nums {
//		update(v, 1)
//
//		if i >= x {
//			update(nums[i-x], -1)
//		}
//
//		if i >= x-1 {
//			// 偶数长度时，取两个中间元素中较小的一个
//			p := countTree.kth(int64((x + 1) / 2))
//			median := int64(unique[p-1])
//
//			// 左侧包含所有 <= median 的元素
//			leftCount := countTree.prefix(p)
//			leftSum := sumTree.prefix(p)
//
//			rightCount := int64(x) - leftCount
//			rightSum := totalSum - leftSum
//
//			cost[i-x+1] =
//				median*leftCount - leftSum +
//					rightSum - median*rightCount
//		}
//	}
//
//	// prev[i]：从 nums[i:] 中选 j-1 个窗口的最小代价
//	// 初始 j=0，不选窗口时，所有位置的代价都是 0
//	prev := make([]int64, n+1)
//	curr := make([]int64, n+1)
//	const inf int64 = 1 << 60
//
//	for j := 1; j <= k; j++ {
//		for i := range curr {
//			curr[i] = inf
//		}
//
//		for i := n - x; i >= 0; i-- {
//			// 不选当前窗口
//			curr[i] = curr[i+1]
//
//			// 选当前窗口
//			if prev[i+x] != inf {
//				take := cost[i] + prev[i+x]
//				if take < curr[i] {
//					curr[i] = take
//				}
//			}
//		}
//
//		prev, curr = curr, prev
//	}
//
//	return prev[0]
//}
