package leecode

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

func Test_flatten(t *testing.T) {
	node := NewTree(1)
	node.LR(2, 5).R(6)
	node.Left.LR(3, 4)
	flatten(node)
	node.Print()

	node = NewTree(6)
	node.L(2).RL(3, 4).L(5)
	sortTree(&node)
	node.Print()
}
