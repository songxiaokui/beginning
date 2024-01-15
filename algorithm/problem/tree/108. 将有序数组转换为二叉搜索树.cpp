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
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        TreeNode* root = nullptr;
        for (auto iter = nums.begin(); iter != nums.end(); iter++)
        {
            cout << "---" << *iter << endl;
            insert(&root, *iter);
        }
        // 遍历
        dfs(root);
        return root;
    }

    void insert(TreeNode** root, int target)
    {
        if (!*root)
        {
            *root = new TreeNode(target, nullptr, nullptr);
            return;
        }
        // 根据当前值查找需要插入的节点
        TreeNode* prev = nullptr;
        TreeNode* current = *root;
        while (current)
        {
            if (current->val == target)
            {
                // 节点存在
                return;
            }
            prev = current;
            if (target > current->val)
            {
                current = current->right;
            } else {
                current = current->left;
            }
        }
        if (!prev)
        {
            return;
        }
        if (target > prev->val)
        {
            prev->right = new TreeNode(target, nullptr, nullptr);
        }  else {
            prev->left = new TreeNode(target, nullptr, nullptr);
        }
    }

    void dfs(TreeNode* root)
    {
        if (!root)
        {
            return;
        }
        dfs(root->left);
        cout << (root)->val << "->";
        dfs(root->right);
    }
    */
    TreeNode* sortedArrayToBST(vector<int>& nums) {
        // 已知是有序数组
        // 直接对半分 然后让列表两侧进行构建
        int mr = nums.size();
        return build(nums, 0, mr-1);
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
};

/*
解题思路:
    首先 题目给出了数组是有序的
    所以直接从数组最中间的元素构建 根
    然后在从数组左侧的数据 构建左子树
    从数组的右侧的数据构建右子树
    还是满足上述条件 取中间的元素为根进行递归构建
    当索引left > right 返回空叶子结点即可
*/