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

/*
解题思路:
    与最大深度相反, 递归处理节点
    如果 节点为空 深度为0
    否则如果左子树不为空 1+左子树的深度递归
    如果右子树不为空 1+右子树的深度递归
    如果 左子树和右子树都是空 则说明是叶子节点 深度为1
    否则 返回1+min(左子树的深度递归, 右子树的深度递归)
*/