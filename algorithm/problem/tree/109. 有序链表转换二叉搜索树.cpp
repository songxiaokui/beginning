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
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    /*
    TreeNode* sortedListToBST(ListNode* head) {
        // 可以采用先把链表转换为数组 然后用有序数组的方案实现
        vector<int> li = convertSlice(head);
        int rl = li.size();
        return build(li, 0, rl-1);

    }

    TreeNode* build(vector<int>& nums, int left, int right)
    {
        if (left > right)
        {
            return nullptr;
        }
        int mid = (left + right) / 2;
        TreeNode* root = new TreeNode(nums[mid], nullptr, nullptr);
        root->left = build(nums, left, mid-1);
        root->right = build(nums, mid+1, right);
        return root;
    }

    vector<int> convertSlice(ListNode* head)
    {
        vector<int> li;
        while(head)
        {
            li.push_back(head->val);
            head = head->next;
        }
        return li;
    }
    */
    TreeNode* sortedListToBST(ListNode* head) {
        // 可以定义一个方法获取链表的中间元素
        return build(head, nullptr);
    }

    // 获取链表的中间元素
    ListNode* getMedian(ListNode* left, ListNode* right)
    {
        if (left == right)
        {
            return left;
        }
        ListNode* fast = left;
        ListNode* slow = left;
        while (fast != right && fast->next != right)
        {
            fast = fast->next->next;
            slow = slow->next;
        }
        return slow;
    }

    TreeNode* build(ListNode* head, ListNode* tail)
    {
        if (head == tail)
        {
            return nullptr;
        }
        // 获取中间节点
        ListNode* mid = getMedian(head, tail);
        // 构造根
        TreeNode* root = new TreeNode(mid->val, nullptr, nullptr);
        // 从链表两侧分别构建
        root->left  = build(head, mid);
        root->right = build(mid->next, tail);
        return root;
    }

};

/*
解题思路:
    首先 要吧有序的链表转换为平衡二叉搜索树
    首先就要确定链表的中间元素 root
    所以 我们可以从给定链表的头部到尾部 获取一个链表的中间节点方法
    然后安装有序数组的方式
    先构建中间节点
    然后从中间节点到尾部构建右子树
    从头节点到中间节点构建左子树
    如果 头和尾相同 则返回为空
*/