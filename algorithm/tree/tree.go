package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}
