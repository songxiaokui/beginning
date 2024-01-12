
// 对称判断 等价判断一个节点的左子树是否等于其节点的右子树同时右子树是否等于左子树

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
    bool isSymmetric(TreeNode* root) {
        if (!root)
        {
            return true;
        }
        return isSeam(root->left, root->right);
    }

    // 从节点拆开,就和判断两个树是否相同
    bool isSeam(TreeNode* p, TreeNode* q)
    {
        // 全假判断
        if (p == nullptr && q == nullptr)
        {
            return true;
        }
        // 一真一假
        if (!(p && q))
        {
            return false;
        }
        // 全为真判断值相同
        if (p->val != q->val)
        {
            return false;
        }
        return isSeam(p->left, q->right) && isSeam(p->right, q->left);
    }
};