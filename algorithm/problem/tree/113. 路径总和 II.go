package tree

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
*/

/*
解题思路:
	使用一个二维数组的地址透传 然后使用一个数组存储当前阶段访问过的元素
	注意 在使用前先拷贝数组
	避免 同数组在递归中出现混乱
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	total := make([][]int, 0)
	current := make([]int, 0)
	check(root, targetSum, &total, current)
	return total
}

func check(root *TreeNode, targetSum int, total *[][]int, current []int) {
	// 此处注意
	current1 := make([]int, 0)
	current1 = append(current1, current...)
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		current1 = append(current1, root.Val)
		*total = append(*total, current1)
		return
	}

	current1 = append(current1, root.Val)
	if root.Left != nil {
		check(root.Left, targetSum-root.Val, total, current1)
	}
	if root.Right != nil {
		check(root.Right, targetSum-root.Val, total, current1)
	}
}
