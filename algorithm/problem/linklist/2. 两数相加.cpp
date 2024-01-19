/*
2. 两数相加
已解答
中等
相关标签
相关企业
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。



示例 1：


输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* head = new ListNode(0, nullptr);
        ListNode* current = head;
        if (!l1 && !l2)
        {
            return head;
        }
        if (!l1)
        {
            return l2;
        }
        if (!l2)
        {
            return l1;
        }
        int carry = 0;
        while (l1 || l2)
        {
            // 只要l1 或者 l2 不同时为 0
            int val = 0;
            int v1 = 0;
            int v2 = 0;
            if (l1)
            {
                v1 = l1->val;
                l1 = l1->next;
            }

            if (l2)
            {
                v2 = l2->val;
                l2 = l2->next;
            }

            // 计算数值与进位
            ListNode* node = new ListNode((v1+v2+carry)%10, nullptr);
            carry = (v1+v2+carry)/10;
            current->next = node;
            current = current->next;
        }

        // 处理最后一进位
        if (carry)
        {
            current->next = new ListNode(carry, nullptr);
        }
        return head->next;
    }
};

// go
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head  *ListNode = &ListNode{
        Val: 0,
        Next: nil,
        }
    var current *ListNode = head
    if l1 == nil && l2 == nil {
        return head
    }

    var carry int = 0
    for (l1 != nil || l2 != nil) {
        var v1 int
        var v2 int
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        // 创建节点
        var node *ListNode = &ListNode{
            Val: (v1+v2+carry) % 10,
            Next: nil,
        }
        carry = (v1+v2+carry) / 10

        current.Next = node
        current = current.Next
    }
    if carry != 0 {
        current.Next = &ListNode{
            Val: carry,
            Next: nil,
        }
    }
    return head.Next
}
*/
