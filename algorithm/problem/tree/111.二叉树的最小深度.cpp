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
    int minDepth(TreeNode* root) {
        if (!root)
        {
            return 0;
        }
        if (!root->left)
        {
            return 1 + minDepth(root->right);
        }
        if (!root->right)
        {
            return 1+minDepth(root->left);
        }
        // 此时到了叶子结点 深度+1
        if (root->left ==nullptr && root->right == nullptr)
        {
            return 1;
        }
        // 每递归一次 深度+1
        return 1+min(minDepth(root->left), minDepth(root->right));
    }

    int min(int a, int b)
    {
        return a < b?a:b;
    }
};