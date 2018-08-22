package leecode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(a *[]int) *ListNode {
	count:=len(*a)
	var parent *ListNode
	node := &ListNode{}
	root := node
	for i := 0; i < count; i++ {
		
		node.Val = (*a)[i]
		if(parent != nil){
			parent.Next = node
		}
		parent = node

		if i < count-1{
			node.Next = &ListNode{}
			node = node.Next
		}
	}
	return root
}

func PrintListNode(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		fmt.Print("->")
		l = l.Next
	}
	fmt.Println()
}

