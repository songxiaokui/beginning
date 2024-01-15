/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* deleteDuplicates(ListNode* head) {
        // 与删除重复元素类似 但是需要记录一下 删除的上一个元素 判断3-4连的情况
        if (!head)
        {
            return head;
        }
        ListNode* dummy_node = new ListNode(0, head);
        ListNode* current = dummy_node;
        while (current->next && current->next->next)
        {
            if (current->next->val == current->next->next->val)
            {
                // 保留当前的值
                int v = current->next->val;
                while (current->next && current->next->val == v)
                {
                    current->next = current->next->next;
                }
            } else
            {
                current = current->next;
            }
        }
        return dummy_node->next;
    }
};