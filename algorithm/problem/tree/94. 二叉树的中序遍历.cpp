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
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        dfs(result, root);
        return result;
    }

    void dfs(vector<int>& f, TreeNode* root)
    {
        if (!root)
        {
            return;
        }
        dfs(f, root->left);
        f.push_back(root->val);
        dfs(f, root->right);
    }
};