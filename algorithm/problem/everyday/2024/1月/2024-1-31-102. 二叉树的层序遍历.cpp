/*
102. 二叉树的层序遍历
中等
相关标签
相关企业
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
*/

class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
        // 定义一个队列处理当前层元素 并将根节点押入 我们使用vector模拟队列实现
        vector<TreeNode*> q1;
        q1.push_back(root);
        // 处理当前层元素
        while (!q1.empty())
        {
            // 定义一个容器处理当前层的值
            vector<int> data;
            // 定义一个零时队列保存下一层的节点
            vector<TreeNode*> next;
            while (!q1.empty())
            {
                // 获取当前层队列的头部元素
                TreeNode* node = q1.front();
                // 删除当前队列的头部元素
                q1.erase(q1.begin());
                if (node)
                {
                    data.push_back(node->val);
                    // 处理左子树
                    if (node->left) next.push_back(node->left);
                    // 处理右子树
                    if (node->right) next.push_back(node->right);
                }
            }
            // 保存当前处理层的数据结果
            if (!data.empty())
            {
                result.push_back(data);
            }
            // 处理下一层数据
            q1 = next;
        }
        return result;
    }
};

/*
解题思路:
    1. 使用一个队列去实现当前层的元素处理
    2. 使用另外一个队列处理当前层的下一层的元素
    3. 当当前层处理完毕 则将当前层的值押入结果队列 并更新处理为下一层队列
*/