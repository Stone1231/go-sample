package leecode

import (
	"fmt"
	"testing"
	. "github.com/sample/dsa"
)

func invertTree(root *TreeNode) *TreeNode {

	if(root==nil){
		return root		
	}

	root.Right, root.Left = root.Left, root.Right
	if	root.Right != nil{
		invertTree(root.Right)
	}
	if	root.Left != nil{
		invertTree(root.Left)
	}
	return root	
}

func Test_invertTree(t *testing.T) {
	node := NewTree(4)
	node.LR(2,7).LR(6,9)
	node.Left.LR(1,3)
	node.Print()
	fmt.Println("=========after=========")
	invertTree(node)
	node.Print()
}