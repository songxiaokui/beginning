package tree

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return build(head, nil)
}

func getMedian(head, tail *ListNode) *ListNode {
	if head == tail {
		return head
	}
	var fast, slow = head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func build(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{
		Val:   mid.Val,
		Left:  nil,
		Right: nil,
	}
	root.Left = build(left, mid)
	root.Right = build(mid.Next, right)
	return root
}
