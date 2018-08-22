package leecode

import (
	. "github.com/sample/dsa"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {

	m := map[int]*ListNode{}

	root := &ListNode{
		Val:  0,
		Next: head,
	}
	parent := root

	for head != nil {
		if _, ok := m[head.Val]; ok {
			parent.Next = head.Next

			removeParent := m[head.Val]
			if removeParent != nil {
				
				if removeParent.Next == parent{
					parent = removeParent
				}

				removeParent.Next = removeParent.Next.Next
			}

			m[head.Val] = nil
		} else {
			m[head.Val] = parent
			parent = head
		}
		head = head.Next
	}

	return root.Next
}

func Test_deleteDuplicates(t *testing.T) {

	l1 := GetListNode(&[]int{1, 2, 3, 3, 4, 4, 5})
	d1 := deleteDuplicates(l1)
	PrintListNode(d1)

	l2 := GetListNode(&[]int{1, 1, 1, 2, 3})
	d2 := deleteDuplicates(l2)
	PrintListNode(d2)
}
