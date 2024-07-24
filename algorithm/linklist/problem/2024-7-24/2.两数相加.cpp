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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* head = new ListNode();
        if (l1 == nullptr && l2 == nullptr) {
            return head->next;
        }
        ListNode* current = head;
        int carry(0);

        while (l1 != nullptr || l2 != nullptr) {
            int v1(0);
            int v2(0);
            ListNode* node = new ListNode();

            if (l1) {
                v1 = l1->val;
                l1 = l1->next;
            }
            if (l2) {
                v2 = l2->val;
                l2 = l2->next;
            }

            node->val = (v1+v2+carry)%10;
            carry = (v1+v2+carry)/10;
            current->next = node;
            current = current->next;
        }
        if (carry != 0) {
            current->next = new ListNode(carry);
        }
        return head->next;
    }
}
