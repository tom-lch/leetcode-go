package SwapNodesinPairs


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var h = &ListNode{}
	h.Next = head
	l := h
	p, q := head, head.Next
	for p != nil && q != nil {
		l.Next = q
		q.Next, p.Next = p, q.Next
		l = p
		if p.Next == nil || p == nil {
			break
		}
		p = p.Next
		q = p.Next
	}
	return h.Next
}

