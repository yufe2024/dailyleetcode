package novemember

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	last := head
	for i := 1; i < k; i++ {
		last = last.Next
	}
	ll := head
	for last.Next != nil {
		ll, last = last.Next, ll.Next
	}
	ss := head
	for i := 1; i < k; i++ {
		ss = ss.Next
	}
	ll.Val, ss.Val = ss.Val, ll.Val
	return head
}
