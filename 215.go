package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	for i := 1; i < k; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

//import (
//"container/heap"
//"fmt"
//)
//
//// 1. 定义一个切片类型的别名
//type IntHeap []int
//
//// 2. 实现 sort.Interface 的三个方法
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆是 <，大顶堆改成 >
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//// 3. 实现 heap.Interface 的 Push 方法（注意：必须是指针接收者）
//func (h *IntHeap) Push(x any) {
//	*h = append(*h, x.(int))
//}
//
//// 4. 实现 heap.Interface 的 Pop 方法（注意：必须是指针接收者）
//func (h *IntHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func main() {
//	h := &IntHeap{2, 1, 5}
//
//	// 5. 初始化堆（这一步会调整切片顺序，使其符合堆特性）
//	heap.Init(h)
//
//	// 6. 插入元素（必须使用 heap.Push，不能用 h.Push）
//	heap.Push(h, 3)
//	heap.Push(h, 0)
//
//	fmt.Println("当前堆顶元素（最小值）:", (*h)[0]) // 0
//
//	// 7. 弹出元素（必须使用 heap.Pop，会触发堆化调整）
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h).(int))
//	}
//	// 输出: 0 1 2 3 5 （从小到大排序输出）
//}
