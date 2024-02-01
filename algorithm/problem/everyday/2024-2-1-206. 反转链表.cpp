/*
206. 反转链表
相关标签
相关企业
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
*/

class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        // 使用一个节点记录先前处理的节点
        ListNode* prev = nullptr;
        ListNode* move = head;
        while (move)
        {
            // 保存move下次要移动的节点
            ListNode* temp = move->next;
            // 反转当前节点
            move->next = prev;
            // 更改prve指向
            prev = move;
            // 更改move指向下一个处理元素
            move = temp;
        }
        return prev;
    }
};

// go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    current := head
    for current != nil {
        temp := current.Next
        current.Next = prev
        prev = current
        current = temp
    }
    return prev
}

// python3
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prve = None
        cur = head
        while (cur != None):
            temp = cur.next
            cur.next=prve
            prve=cur
            cur=temp
        return prve