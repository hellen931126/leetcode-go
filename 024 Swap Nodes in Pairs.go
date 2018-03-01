func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	second := head.Next
	if second == nil {
		return head
	}
	head.Next = swapPairs(second.Next)
	second.Next = head
	return second
}