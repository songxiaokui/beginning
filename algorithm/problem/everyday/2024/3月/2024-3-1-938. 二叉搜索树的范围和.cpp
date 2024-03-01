/*
938. 二叉搜索树的范围和
简单
相关标签
相关企业
给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
输出：32
*/

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
    int m_low;
    int m_high;
    int result;
    int rangeSumBST(TreeNode* root, int low, int high) {
        this->m_high = high;
        this->m_low = low;
        bfs(root);
        return result;
    }

    void bfs(TreeNode* root)
    {
        if (!root)
        {
            return ;
        }
        queue<TreeNode*> mq;
        mq.push(root);
        while (!mq.empty())
        {
            // 获取头部元素
            TreeNode* node = mq.front();
            // 删除头部元素
            mq.pop();
            // 处理元素
            if (node->val <= this->m_high && node->val >= m_low)
            {
                this->result += node->val;
            }
            if (node->left) mq.push(node->left);
            if (node->right) mq.push(node->right);
        }
    }
};

// Go
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
    var result int
    dfs(root, low, high, &result)
    return result
}

func dfs(root *TreeNode, low,high int, result *int) {
    if root == nil {
        return
    }
    if root.Val <= high && root.Val >= low {
        *result += root.Val
    }
    dfs(root.Left, low, high, result)
    dfs(root.Right, low, high, result)
    return
}