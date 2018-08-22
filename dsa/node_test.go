package dsa
import "testing"

func Test_listNode(t *testing.T) {
	list := GetListNode(&[]int{1, 2, 3, 4, 5})
	PrintListNode(list)
}

func Test_treeNode(t *testing.T) {
	node := NewTree(0)
	node.LR(12,3).RL(4,19).LR(15,6).LR(17,8)
	node.Right.Left.Left.LR(11,6)
	node.Left.LR(7,13).LR(2,22)
	node.Print()
}