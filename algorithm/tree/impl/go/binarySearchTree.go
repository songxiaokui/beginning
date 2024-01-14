package _go

import "fmt"

// BinarySearchTree 二叉搜索树
type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

func (bst *BinarySearchTree) Insert(value int) {
	currentNode := bst.root
	if bst.root == nil {
		bst.root = NewTreeNode(value)
		return
	}
	var prvNode *TreeNode = nil // 用来保存最后需要处理的节点
	for currentNode != nil {
		if currentNode.Value == value {
			// 插入的节点等于当前节点
			// 节点存在
			return
		}
		prvNode = currentNode
		if currentNode.Value > value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	// 循环结束 找到插入的节点 判断插入左侧还是右侧
	if prvNode.Value > value {
		prvNode.Left = NewTreeNode(value)
	} else {
		prvNode.Right = NewTreeNode(value)
	}
}

func (bst *BinarySearchTree) Search(value int) *TreeNode {
	if bst.root == nil {
		return nil
	}
	currentNode := bst.root
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		if currentNode.Value > value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return nil
}

func (bst *BinarySearchTree) Delete(value int) {
	// 保存当前节点的先驱节点
	var prvNode *TreeNode = nil
	var currentNode *TreeNode = bst.root
	if currentNode == nil {
		return
	}
	for currentNode != nil {
		if currentNode.Value == value {
			break
		}
		prvNode = currentNode
		if currentNode.Value > value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	// 判断是否差到需要删除的节点
	if currentNode == nil {
		// 无节点
		return
	}

	// 情况1: 删除的节点为叶子节点 节点度为0
	// 情况2: 删除的节点有一个子节点 节点度为1
	if currentNode.Left == nil || currentNode.Right == nil {
		var ch *TreeNode = nil
		if currentNode.Left != nil {
			ch = currentNode.Left
		} else {
			ch = currentNode.Right
		}
		// 删除当前节点
		if currentNode != bst.root {
			if prvNode.Left == currentNode {
				prvNode.Left = ch
			} else {
				prvNode.Right = ch
			}
		} else {
			currentNode = ch
		}
	} else {
		// 情况3: 删除的节点有2个子节点 节点度为2
		// 将当前删除的节点的值 替换为中序遍历的后继节点
		// 递归进行后继节点删除
		temp := currentNode.Right
		for temp.Left != nil {
			temp = temp.Left
		}

		// 此时找到当前删除节点的后继节点
		// 删除该节点
		bst.Delete(temp.Value)
		// 把当前节点删除节点的值修改
		currentNode.Value = temp.Value
	}
}

func (bst *BinarySearchTree) Print() {
	// 使用中序遍历输出
	bst.inorderDFS(bst.root)
}

func (bst *BinarySearchTree) inorderDFS(node *TreeNode) {
	if node == nil {
		return
	}
	// 先处理左子树
	bst.inorderDFS(node.Left)
	// 处理当前节点
	fmt.Printf("%d -> ", node.Value)
	// 处理右子树
	bst.inorderDFS(node.Right)
}
