class Solution {
public:
    bool hasPathSum(TreeNode* root, int targetSum) {
        if (root == nullptr)
        {
            return 0;
        }
        if (root->left == nullptr && root->right == nullptr)
        {
            return targetSum - root->val == 0;
        }
        // 从左子树和右子树上分别查找
        return hasPathSum(root->left, targetSum-root->val)||hasPathSum(root->right, targetSum-root->val);
    }

};

/*
解题思路:
    每次递归进入，判断 是否 左、右节点为空 如果为空且目标值等于当前节点的值 则可以找到当前路径总和
    否则
    从左子树 或者 右子树递归判断是否满足上述条件
    满足上述条件一个即可
*/