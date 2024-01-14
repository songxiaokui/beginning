package _go

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	// 创建
	tree := NewBinarySearchTree()
	// 插入
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)

	// 打印
	tree.Print()

	// 删除一个节点为3
	tree.Delete(3)
	fmt.Println("\n删除3之后的二叉搜索树:")
	tree.Print()

	// 插入3和6 1个节点删除
	tree.Insert(3)
	tree.Insert(6)
	fmt.Println("\n插入3和6之后的二叉搜索树:")
	tree.Print()

	// 删除7节点 单节点删除
	fmt.Println("\n删除7之后的二叉搜索树:")
	tree.Delete(7)
	tree.Print()

	// 还原数据7 删除2个度的5
	fmt.Println("\n删除5之后的二叉搜索树:")
	tree.Insert(7)
	tree.Delete(5)
	tree.Print()

	// 查找
	t.Log(tree.Search(8) == nil)
	t.Log(tree.Search(15) != nil)
}
