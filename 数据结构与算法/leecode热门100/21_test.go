package leecode热门100

import "testing"

func Test_MergeTwoLists(t *testing.T) {
	var l1 = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	var l2 = &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	MergeTwoLists(l1, l2)
}
