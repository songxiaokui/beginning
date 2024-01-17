// 暴力法 将数据中序遍历,然后进行比较
class Solution {
public:
    bool isValidBST(TreeNode* root) {
        // 中序遍历 判断是否大于当前
        vector<int> li;
        dfs(root, li);
        int l = li.size();
        for (int i = 0; i < l-1; i++)
        {
            if (li[i] >= li[i+1])
            {
                return false;
            }
        }
        return true;
    }

    void dfs(TreeNode* root, vector<int>& li)
    {
        if (!root)
        {
            return;
        }
        // 左
        dfs(root->left, li);
        // 数据处理
        li.push_back(root->val);
        // 右
        dfs(root->right, li);
    }
};

// 深度优先 (记录一个最小值 如果存在节点比最小值大 则报错)
class Solution {
public:
    long long min_val = LONG_MIN;
public:
    bool isValidBST(TreeNode* root)
    {
       // 可以把根节点记录
       if (!root)
       {
           return true;
       }
       // 递归时只要左节点值大于根 就是false
       // 最后要满足 左 右 都满足
       return dfs(root);

    }

    bool dfs(TreeNode* root)
    {
        if (!root)
        {
            return true;
        }
        // 左
        bool left = dfs(root->left);
        // 中
        if (min_val < root->val)
        {
            min_val = root->val;
        } else {
            return false;
        }
        // 右
        bool right = dfs(root->right);
        return right&&left;
    }
};