package easy

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	tmp := &ListNode{}
	head := tmp

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}
		head = head.Next
		if l1 == nil {
			head.Next = l2
		}
		if l2 == nil {
			head.Next = l1
		}
	}
	return tmp.Next
}
