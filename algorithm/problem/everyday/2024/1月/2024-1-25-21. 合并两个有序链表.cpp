/*
21. 合并两个有序链表
已解答
简单
相关标签
相关企业
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
*/

ListNode *mergeTwoLists(ListNode *list1, ListNode *list2) {
        ListNode *dummyNode = new ListNode(0, nullptr);
        ListNode *result = dummyNode;

        while (list1 && list2) {
            // 比较值 确定指向
            if (list1->val <= list2->val) {
                dummyNode->next = list1;
                list1 = list1->next;
            } else {
                dummyNode->next = list2;
                list2 = list2->next;
            }
            dummyNode = dummyNode->next;
        }

        if (list1) {
            dummyNode->next = list1;
        }
        if (list2) {
            dummyNode->next = list2;
        }

        return result->next;
    }

// go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyNode := &ListNode{
        Val:0,
        Next: nil,
    }

    var currentNode *ListNode = dummyNode

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            currentNode.Next = list1
            list1 = list1.Next
        } else {
            currentNode.Next = list2
            list2 = list2.Next
        }
        currentNode = currentNode.Next
    }
    if list1 != nil {
        currentNode.Next = list1
    }
    if list2 != nil {
        currentNode.Next = list2
    }
    return dummyNode.Next
}

// python3
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummyNode = ListNode(0,None)
        currentNode = dummyNode
        while list1 != None and list2 != None:
            if list1.val <= list2.val:
                currentNode.next = list1
                list1 = list1.next
            else:
                currentNode.next = list2
                list2 = list2.next
            currentNode = currentNode.next
        if list1 != None:
            currentNode.next = list1
        if list2 != None:
            currentNode.next = list2
        return dummyNode.next