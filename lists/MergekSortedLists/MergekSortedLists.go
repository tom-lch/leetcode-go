package MergekSortedLists


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

// 使用分治法解决多个有序数组合并成一个有序数组

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	mid := len(lists)/2
	l1 := make([]*ListNode, mid)
	copy(l1, lists)
	l2 :=  make([]*ListNode, len(lists)-mid)
	copy(l2, lists[mid:])
	return mergeTwoLists(mergeKLists(l1), mergeKLists(l2))
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{}
	head := node
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}
