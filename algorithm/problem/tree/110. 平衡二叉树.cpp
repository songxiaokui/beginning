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
    bool isBalanced(TreeNode* root) {
        if (!root)
        {
            return true;
        }
        return isBalanced(root->left) && isBalanced(root->right) && abs(height(root->left) - height(root->right)) <= 1;
    }

    int max(int a, int b)
    {
        return a > b ? a : b;
    }

    // 获取树的最大高度
    int height(TreeNode* root)
    {
        if (!root)
        {
            return 0;
        }
        return max(height(root->left), height(root->right)) + 1;
    }
};

/*
解题思路:
    判断一个树是否是平衡二叉树
    首先 节点的左子树与节点的右子树高度差的绝对值<=1
    同时 对于任何一个节点 都要满足上述条件
*/