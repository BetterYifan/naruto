package easy

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	var prev = &ListNode{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		// 反转前半部分链表
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for slow != nil && prev != nil {
		if slow.Val != prev.Val {
			return false
		}
		// slow继续往后遍历
		slow = slow.Next
		prev = prev.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	return reverseNode(nil, head)
}

func reverseNode(prev, curr *ListNode) *ListNode {
	if curr == nil {
		return prev
	}
	next := curr.Next
	curr.Next = prev
	return reverseNode(curr, next)
}
