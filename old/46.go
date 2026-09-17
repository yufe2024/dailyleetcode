package old

func permute(nums []int) [][]int {
	var ans [][]int
	var used = make([]bool, len(nums))
	var path []int
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				path = append(path, nums[i])

				dfs()

				path = path[:len(path)-1]
				used[i] = false
			} else {
				continue
			}

		}
	}
	dfs()
	return ans
}

//package main
//
//import "fmt"
//
//func permute(nums []int) [][]int {
//	var ans [][]int
//	var path []int
//	// used 数组用来标记当前数字是否已经在路径中
//	used := make([]bool, len(nums))
//
//	// 将 dfs 定义在内部，可以直接闭包引用外部的 ans, path, used，避免全局变量
//	var dfs func()
//	dfs = func() {
//		// 结束条件：当前路径长度等于数组长度，说明找到了一个完整排列
//		if len(path) == len(nums) {
//			// CRITICAL: Go 的切片是引用传递，必须 copy 一份，否则后续修改会影响 ans 里的结果
//			temp := make([]int, len(path))
//			copy(temp, path)
//			ans = append(ans, temp)
//			return
//		}
//
//		for i := 0; i < len(nums); i++ {
//			// 如果这个数字已经用过了，跳过
//			if used[i] {
//				continue
//			}
//
//			// 做选择
//			path = append(path, nums[i])
//			used[i] = true
//
//			// 递归
//			dfs()
//
//			// 撤销选择（回溯的核心）
//			path = path[:len(path)-1]
//			used[i] = false
//		}
//	}
//
//	dfs()
//	return ans
//}
//
//func main() {
//	nums := []int{1, 2, 3}
//	result := permute(nums)
//	fmt.Println(result)
//	// 输出: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
//}
