/*
993. 二叉树的堂兄弟节点
简单
相关标签
相关企业
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。

示例 1：
输入：root = [1,2,3,4], x = 4, y = 3
输出：false
*/

class Solution {
public:
    // 解题思路
    // 找到x和y节点的深度并记录x和y的父亲节点

    // x的基本信息
    int x;
    TreeNode* xp;
    int x_depth;
    bool x_found = false;

    // y的基本信息
    int y;
    TreeNode* yp;
    int y_depth;
    bool y_found = false;

    void dfs(TreeNode* node, int depth, TreeNode* parent)
    {
        if (node == nullptr)
        {
            return;
        }
        if (node->val == this->x)
        {
            // 使用tie和tuple 将数据解构
            tie(xp, x_depth, x_found) = tuple(parent, depth, true);
        }
        else if (node->val == this->y)
        {
            tie(yp, y_depth, y_found) = tuple(parent, depth, true);
        }
        // 如果两个元素提前找到 则退出
        if (x_found && y_found)
        {
            return;
        }

        // 否则继续递归左子树查找
        dfs(node->left, depth+1, node);

        // 继续判断是否提前找到
        if (x_found && y_found)
        {
            return;
        }

        // 递归右子树查找
        dfs(node->right, depth+1, node);
    }

    bool isCousins(TreeNode* root, int x, int y) {
        this->x = x;
        this->y = y;
        dfs(root, 0, root);
        return x_depth==y_depth && xp != yp;
    }
};

// go
/*
type Result struct {
     Data int
     Depth int
     Parent *TreeNode
     IsFound bool
 }

func dfs(root *TreeNode, depth int, parent *TreeNode, datax *Result, datay *Result) {
    if root == nil {
        return
    }
    if root.Val == datax.Data {
        datax.Depth = depth
        datax.Parent = parent
        datax.IsFound = true
    } else if root.Val == datay.Data {
        datay.Depth = depth
        datay.Parent = parent
        datay.IsFound = true
    }

    if datax.IsFound && datay.IsFound {
        return
    }

    dfs(root.Left, depth+1, root, datax, datay)

    if datax.IsFound && datay.IsFound {
        return
    }

     dfs(root.Right, depth+1, root, datax, datay)
}

func isCousins(root *TreeNode, x int, y int) bool {
    xi := Result{}
    xi.Data = x

    yi := Result{}
    yi.Data = y

    dfs(root, 0, root, &xi, &yi)

    return xi.Depth==yi.Depth && xi.Parent != yi.Parent
}
*/