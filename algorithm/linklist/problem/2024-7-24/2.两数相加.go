package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	result := &ListNode{
		Val:  0,
		Next: nil,
	}
	current := result
	carry := 0
	for l1 != nil || l2 != nil {
		var (
			v1   int
			v2   int
			node = new(ListNode)
		)
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		node.Next = nil
		node.Val = (v1 + v2 + carry) % 10

		current.Next = node
		current = current.Next

		carry = (v1 + v2 + carry) / 10
	}

	if carry != 0 {
		current.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return result.Next
}
