package everyday

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 定义哑节点
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	// 定义快慢指针
	var fast, slow *ListNode = dummyNode, dummyNode

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyNode.Next
}
