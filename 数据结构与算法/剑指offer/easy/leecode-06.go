package easy

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	l := list.New()
	for head != nil {
		l.PushFront(head.Val)
		head = head.Next
	}
	res := make([]int, 0)
	for l.Len() > 0 {
		res = append(res, l.Front().Value.(int))
		l.Remove(l.Front())
	}
	return res
}
