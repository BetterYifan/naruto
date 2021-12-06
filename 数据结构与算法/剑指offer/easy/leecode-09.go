package easy

import "container/list"

type CQueue struct {
	list1, list2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		list1: list.New(),
		list2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.list1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.list2.Len() == 0 {
		for this.list1.Len() != 0 {
			this.list2.PushBack(this.list1.Remove(this.list1.Back()))
		}
	}
	if this.list2.Len() != 0 {
		val := this.list2.Back()
		this.list2.Remove(val)
		return val.Value.(int)
	}
	return -1
}
