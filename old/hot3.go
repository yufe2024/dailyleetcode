package old

func lengthOfLongestSubstring(s string) int {
	ans := 0
	an := 1
	a := []byte(s)
	left := 0
	for ; left < len(a); left++ {
		right := left + 1
		set := make(map[byte]int)
		set[a[left]] = 1
		for right < len(a) && set[a[right]] == 0 {
			set[a[right]] = 1
			right++
		}
		an = right - left
		if an > ans {
			ans = an
		}
	}
	return ans
}

//func lengthOfLongestSubstring(s string) int {
//	ans := 0
//	set := make(map[byte]bool) // 复用同一个 map
//	left := 0
//
//	// right 一直向右移动，不回头
//	for right := 0; right < len(s); right++ {
//		// 如果遇到了重复字符，说明左边界 left 需要收缩
//		// 边收缩边把离开窗口的字符从 map 中删掉，直到移除了重复的那一个
//		for set[s[right]] {
//			delete(set, s[left])
//			left++
//		}
//
//		// 将当前字符加入集合
//		set[s[right]] = true
//
//		// 此时整个窗口内绝对没有重复字符，且右边界就在当前字符上
//		// 因为 left 和 right 都是合法的闭区间，所以长度是 right - left + 1
//		if right-left+1 > ans {
//			ans = right - left + 1
//		}
//	}
//	return ans
//}
