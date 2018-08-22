package dsa

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(a *[]int) *ListNode {
	count := len(*a)
	var parent *ListNode
	node := &ListNode{}
	root := node
	for i := 0; i < count; i++ {

		node.Val = (*a)[i]
		if parent != nil {
			parent.Next = node
		}
		parent = node

		if i < count-1 {
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(n int) *TreeNode {
	node := &TreeNode{
		Val : n,		
	}
	return node
}

func (self *TreeNode) R(n int) *TreeNode{ 
	if(self.Right ==nil){
		self.Right = NewTree(n)	
	} else {
		self.Right.Val = n
	}
	return self.Right
}

func (self *TreeNode) L(n int) *TreeNode{ 
	if(self.Left ==nil){
		self.Left = NewTree(n)	
	} else {
		self.Left.Val = n
	}
	return self.Left	
}

func (self *TreeNode) LR(l int,r int) *TreeNode{ 
	self.L(l)
	return self.R(r)	
}

func (self *TreeNode) RL(l int,r int) *TreeNode{ 
	self.R(r)
	return self.L(l)	
}

func printTreeNode(p *TreeNode, indent int){
    if(p != nil) {
		if(p.Left != nil) {
			printTreeNode(p.Left, indent+4)}
		if(p.Right != nil) {
			printTreeNode(p.Right, indent+4)}
        if (indent > 0) {
			fmt.Print(strings.Repeat(" ", indent))
            //std::cout << std::setw(indent) << ' ';
		}
		fmt.Println(p.Val)
    }
}

func (self *TreeNode) Print(){
	printTreeNode(self, 0)
}