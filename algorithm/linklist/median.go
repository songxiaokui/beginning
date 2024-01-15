package linklist

// GetMedianElement 获取链表的中位元素
func GetMedianElement(node *LinkNode) *LinkNode {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		return node
	}

	// 使用快慢指针,并定义一个哑指针在链表的头部
	dummy := &LinkNode{
		Val:  0,
		Next: node,
	}
	var slow, fast = dummy, dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func GetMedianElementII(head *LinkNode, tail *LinkNode) *LinkNode {
	if head == tail {
		return head
	}
	fast := head
	slow := head

	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
