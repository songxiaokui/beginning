package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var dummy *ListNode = &ListNode{
		Val:  0,
		Next: head,
	}
	var current *ListNode = dummy
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			// 说明遇到开始重复的元素 统一处理掉 直到没
			val := current.Next.Val
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}
	return dummy.Next
}

/*
解题思路:
	需要定义一个哑节点用来保存链表的头信息
	然后使用一个当前节点指向哑节点
	哑节点的next为处理节点的头节点
	当cur 节点next == cur.next.next 表示遇到重复元素
	然后记录重复元素 一致循环处理 知道遇到不同元素
	然后 将当前元素指向cur.next
	最后返回 dummy.next即可
*/
