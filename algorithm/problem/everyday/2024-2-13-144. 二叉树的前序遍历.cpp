/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
*/

class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        post(root, result);
        return result;
    }

    void post(TreeNode* node, vector<int>& result)
    {
        if (!node)
        {
            return;
        }
        result.push_back(node->val);
        post(node->left, result);
        post(node->right, result);
    }
};