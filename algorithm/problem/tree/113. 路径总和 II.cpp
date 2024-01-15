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
    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> init;
        vector<int> current;
        check(root, targetSum, init, current);
        return init;
    }

    void check(TreeNode* root, int targetSum, vector<vector<int>>& total, vector<int> current)
    {
        // 创建一个数组 将引用的值取出
        vector<int> newAarray(current);
        if (!root)
        {
            return;
        }
        if (!root->left && !root->right && targetSum - root->val == 0)
        {
            newAarray.push_back(root->val);
            total.push_back(newAarray);
            return;
        }
        newAarray.push_back(root->val);
        // 处理左子树
        if (root->left)
        {
            check(root->left, targetSum-root->val, total, newAarray);
        }

        // 处理右子树
        if (root->right)
        {
            check(root->right, targetSum-root->val, total, newAarray);
        }

        return;
    }
};

/*
注意:
    total需要改变需要传递引用
    current防止左右子树同时修改 进行值传递
*/