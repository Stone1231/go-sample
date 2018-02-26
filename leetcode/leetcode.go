package main

import (
	"strconv"
)

func main() {
	addTwoNumbersTest()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbersTest() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	list := addTwoNumbers(&l1, &l2)

	for list != nil {
		println(list.Val)
		println("->")
		list = list.Next
	}

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := sumList(l1, 1)
	sum += sumList(l2, 1)
	println(sum)

	list := numToListNode(sum)
	return list
}
func sumList(list *ListNode, num int) int {

	if list.Next != nil {
		return list.Val*num + sumList(list.Next, num*10)
	} else {
		return list.Val * num
	}
}
func numToListNode(sum int) *ListNode {
	//str := string(sum)
	str := strconv.Itoa(sum)

	count := len(str)

	list := &ListNode{}
	for index := 0; index < count; index++ {
		_list := list
		_list.Val = int(str[index] - '0') //轉int
		if index < count-1 {
			list = &ListNode{}
			list.Next = _list
		}
	}

	return list
}
