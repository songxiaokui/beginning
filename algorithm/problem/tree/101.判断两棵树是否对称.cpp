
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

/*
解题思路:
    判断一个树是否对称，除了判断该节点的左右子节点的值是否相同 同时还需要判断左右子树是否同时满足上述条件
    两颗树为空 则对称
    如果一个为空 一个不为空 则不对称
    如果两个节点的值不相等 不对称
    递归调用 节点的左与另外一个节点的右 节点的右 与另外一个节点的左 同时满足对称方为对称
*/