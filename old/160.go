package old

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for headA != nil {
		headA.Val = -1
		headA = headA.Next
	}
	for headB != nil {
		if headB.Val == -1 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
