package main

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var currentNode *ListNode = head
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}
