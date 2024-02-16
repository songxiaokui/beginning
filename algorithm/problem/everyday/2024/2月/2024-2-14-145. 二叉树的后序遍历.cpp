/*
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历
*/
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> result;
        post(root, result);
        return result;
    }
    void post(TreeNode* root, vector<int>& result)
    {
        if (!root)
        {
            return;
        }
        post(root->left, result);
        post(root->right, result);
        result.push_back(root->val);
    }
};

// go
func postorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    post(root, &result)
    return result
}

func post(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    post(root.Left, result)
    post(root.Right, result)
    *result = append(*result, root.Val)
}
