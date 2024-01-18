/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    bool hasCycle(ListNode *head) {

        if (!head || !head->next)
        {
            return false;
        }
        ListNode* slow = head;
        ListNode* fast = head->next;

        while (slow != fast)
        {

            if (!fast || !fast->next)
            {
                return false;
            }
            slow = slow->next;
            fast = fast->next->next;

        }
        return true;
    }
};

/*
使用快慢指针
    当快慢指针相遇 则有环 否则 当快指针为空 或者快指针的 next 为空 则无环
*/

// go实现
/*
func hasCycle(head *ListNode) bool {
    // 快慢指针
    if head == nil || head.Next == nil {
        return false
    }
    slow, fast := head, head.Next
    for (slow != fast) {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}
*/