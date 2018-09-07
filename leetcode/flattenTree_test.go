package leetcode

import (
	"testing"

	. "github.com/sample/dsa"
)

func sortNode(root *TreeNode, node *TreeNode) {

	if node.Left != nil {
		sortNode(root, node.Left)
		node.Left = nil
	}

	if node.Right != nil {
		sortNode(root, node.Right)
		node.Right = nil
	}

	current := root

	if current.Right == nil {
		current.Right = node
	} else {
		for current.Right != nil &&
			node.Val > current.Right.Val {
			current = current.Right
		}
		// node.Right = current.Right
		// current.Right = node
		temp := current.Right
		node.Right = temp
		current.Right = node
	}
}

func sortTree(root **TreeNode) {

	newRoot := &TreeNode{
		Left: *root,
	}

	if newRoot.Left != nil {
		sortNode(newRoot, newRoot.Left)
	}

	*root = newRoot.Right

	//只是複製過來的指標,可以透過指標修改原來的值
	//func flatten1(root *TreeNode) {

	//root = newRoot.Right 這樣更改複製的root,不會影響原來的
}

func flatten(root *TreeNode) {

	if (root == nil) ||
		(root.Left == nil && root.Right == nil) {
		return
	}

	if root.Left != nil {
		flatten(root.Left)
	}

	if root.Right != nil {
		flatten(root.Right)
	}

	tmpRight := root.Right
	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmpRight
}

func flattenMy(root *TreeNode) {
	if (root == nil) || (root.Left == nil && root.Right == nil) {
		return
	}

	if root.Left != nil {
		flattenMy(root.Left)
	}

	if root.Right != nil {
		flattenMy(root.Right)
	}
	
	left := root.Left
	if left != nil {
		right := root.Right
		root.Right = left
		root.Left = nil
		for left.Right != nil {
			left = left.Right
		}
		left.Right = right
	}
}

func Test_flatten(t *testing.T) {
	node := NewTree(1)
	node.LR(2, 5).R(6)
	node.Left.LR(3, 4)
	flatten(node)
	node.Print()
}

func Test_flattenMy(t *testing.T) {
	node := NewTree(1)
	node.LR(2, 5).R(6)
	node.Left.LR(3, 4)
	flattenMy(node)
	node.Print()
}

func Test_sortTree(t *testing.T) {
	node := NewTree(6)
	node.L(2).RL(3, 4).L(5)
	sortTree(&node)
	node.Print()
}
