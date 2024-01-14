package tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}
	remain := targetSum - root.Val
	return hasPathSum(root.Left, remain) || hasPathSum(root.Right, remain)
}
