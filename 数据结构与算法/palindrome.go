package 数据结构与算法

func isPalindrome(head *ListNode) bool {
	tail := reverse(findMiddle(head))

	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head.Next
	head.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = head
		head = cur
		cur = next
	}
	return head
}
