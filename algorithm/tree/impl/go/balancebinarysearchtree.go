package _go

type BBSTNode struct {
	Value  any       // 值
	Left   *BBSTNode // 左子树
	Right  *BBSTNode // 右子树
	height int       // 节点的高度
}

// getHeight 获取节点的高度
func (bbst *BBSTNode) getHeight(node *BBSTNode) int {
	if bbst != nil {
		return bbst.height
	}
	return -1
}

// updateHeight 更新节点的高度
func (bbst *BBSTNode) updateHeight() {
	lh := bbst.getHeight(bbst.Left)
	rh := bbst.getHeight(bbst.Right)
	if lh > rh {
		bbst.height = lh + 1
	} else {
		bbst.height = rh + 1
	}
}

// calHeight 计算节点的高度
func (bbst *BBSTNode) calHeight(root *BBSTNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(bbst.calHeight(root.Left), bbst.calHeight(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getBalanceFactor 获取该节点的平衡因子
func (bbst *BBSTNode) getBalanceFactor() int {
	// 递归获取该节点的左子树的最大高度
	// 递归获取该节点的右子树的最大高度
	// 返回lheight-rheight
	if bbst == nil {
		return 0
	}

	return bbst.getHeight(bbst.Left) - bbst.getHeight(bbst.Right)
}
