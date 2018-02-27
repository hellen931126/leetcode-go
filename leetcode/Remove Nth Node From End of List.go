func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	q, p := head, head
	for n > 0 {
		p = p.Next
		n -= 1
	}
	if p == nil {
		return head.Next
	}
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}