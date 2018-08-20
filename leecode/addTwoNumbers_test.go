package leecode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	sum := sumList(l1, 1)
// 	sum += sumList(l2, 1)
// 	println(sum)

// 	list := numToListNode(sum)
// 	return list
// }
// func sumList(list *ListNode, num int) int {

// 	if list.Next != nil {
// 		return list.Val*num + sumList(list.Next, num*10)
// 	} else {
// 		return list.Val * num
// 	}
// }
// func numToListNode(sum int) *ListNode {
// 	//str := string(sum)
// 	str := strconv.Itoa(sum)

// 	count := len(str)

// 	list := &ListNode{}
// 	for index := 0; index < count; index++ {
// 		_list := list
// 		_list.Val = int(str[index] - '0') //轉int
// 		if index < count-1 {
// 			list = &ListNode{}
// 			list.Next = _list
// 		}
// 	}

// 	return list
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := &ListNode{
		Val: 0,
	}

	lists := []*ListNode{l1, l2}
	sumList(sum, lists)

	return sum
}

func sumList(sum *ListNode, lists []*ListNode) {

	hasNext := false
	count := len(lists)
	for index := 0; index < count; index++ {
		list := lists[index]
		hasList := list != nil
		if hasList {
			sum.Val += list.Val
		}

		hasList = hasList && list.Next != nil
		if hasList {
			lists[index] = list.Next
			hasNext = true
		} else {
			lists[index] = nil
		}
	}

	nextVal := sum.Val / 10
	sum.Val = sum.Val % 10

	if hasNext || nextVal > 0 {
		sum.Next = &ListNode{
			Val: nextVal,
		}
		sumList(sum.Next, lists)
	}
}

func Test_addTwoNumbers(t *testing.T) {
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

	// l1 = ListNode{
	// 	Val: 5,
	// }

	// l2 = ListNode{
	// 	Val: 5,
	// }

	list := addTwoNumbers(&l1, &l2)

	for list != nil {
		println(list.Val)
		println("->")
		list = list.Next
	}
}
