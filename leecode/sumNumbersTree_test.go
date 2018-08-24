package leecode

import (
	"fmt"
	. "github.com/sample/dsa"
	"testing"
)

func sumNumbersDFS(root *TreeNode, total int) int {

	if root == nil {
		return 0
	}

	total = total*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return total
	}

	return sumNumbersDFS(root.Left, total) + sumNumbersDFS(root.Right, total)
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersDFS(root, 0)
}

func Test_sumNumbers(t *testing.T) {
	node := NewTree(1)
	node.LR(2, 3)
	fmt.Println(sumNumbers(node))

	node = NewTree(4)
	node.RL(9, 0).LR(5, 1)
	fmt.Println(sumNumbers(node))
}
