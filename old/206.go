package old

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	headB := head.Next
	head.Next = nil
	for head != nil && head.Next != nil {

		head.Next.Next = head
		head.Next = headB
		head = headB
	}
	return headB
}
