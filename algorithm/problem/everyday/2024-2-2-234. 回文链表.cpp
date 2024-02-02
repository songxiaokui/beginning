/*
234. 回文链表
简单
相关标签
相关企业
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
输入：head = [1,2,2,1]
输出：true
*/

class Solution {
public:
    bool isPalindrome(ListNode* head) {
        if (!head)
        {
            return true;
        }
        vector<int> li;
        while (head)
        {
            li.push_back(head->val);
            head = head->next;
        }

        int l = li.size();
        int left=0, right=l-1;
        while (left < right)
        {
            if (li[left] != li[right])
            {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
};