package 数据结构与算法

type ListNode struct {
	Val  int
	Next *ListNode
}

// 移除无序链表中重复的节点
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var tmp = make(map[int]struct{})
	slow := head
	fast := head.Next
	tmp[slow.Val] = struct{}{}
	for fast != nil {
		if _, ok := tmp[fast.Val]; !ok {
			tmp[fast.Val] = struct{}{}
			slow = fast
			fast = fast.Next
		} else {
			fast = fast.Next
		}
		slow.Next = fast
	}
	return head
}
