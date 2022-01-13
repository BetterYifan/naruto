package doublelist

import (
	"fmt"
	"testing"
)

var list = New()

func TestList_PushBack(t *testing.T) {
	t.Run("list", func(t *testing.T) {
		list.PushBack(1)
		list.PushBack(2)
		list.PushBack(3)

		if list.root.next.val.(int) != 1 {
			t.Error("error")
		}
		//fmt.Println(list.root.next.val)
		fmt.Println(list.root.next.next.val)
		fmt.Println(list.root.next.next.next.val)
	})

}
