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
    void flatten(TreeNode* root) {
         vector<TreeNode*> li = scan(root);
         int l = li.size();
         for (int i = 0; i < l-1; i++)
         {
            root->right = li[i+1];
            root->left = nullptr;
            root = root->right;
         }
    }

    vector<TreeNode*> scan(TreeNode* root)
    {
        vector<TreeNode*> li;
        dfs(root, li);
        return li;
    }
    void dfs(TreeNode* root, vector<TreeNode*>& li)
    {
        if (!root)
        {
            return;
        }

        TreeNode* t = new TreeNode(root->val);
        li.push_back(t);
        dfs(root->left, li);
        dfs(root->right, li);
    }

};

/*
解题思路:
    先序遍历获取节点列表
    然后从root的right开始赋值节点
    把left置为nullptr
*/