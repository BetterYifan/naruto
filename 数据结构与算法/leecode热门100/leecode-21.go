package leecode热门100

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	PreHead := &ListNode{0, nil}
	head := PreHead
	for {
		if l1 == nil {
			break
		}
		if l2 == nil {
			break
		}
		if l1.Val <= l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		head = head.Next
	}
	if l1 == nil {
		head.Next = l2
	}
	if l2 == nil {
		head.Next = l1
	}
	return PreHead.Next
}
