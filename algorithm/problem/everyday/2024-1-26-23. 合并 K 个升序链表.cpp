/*
23. 合并 K 个升序链表
已解答
困难
相关标签
相关企业
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。
示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

/*
解题思路
    1. 定义一个从列表中获取最小元素所在索引的位置
    2. 使用哑节点指向该列表索引对应的链表
    3. 更新该索引对应的链表的头节点
    4. 判断最后列表中的所有元素是否处理为空
*/

class Solution {
public:
    // 判断是否都为空
    bool getNull(vector<ListNode*>& lists)
    {
        bool status = false;
        for (auto iter = lists.begin(); iter != lists.end(); iter++)
        {
            if (*iter)
            {
                status = true;
            }
        }
        return status;
    }

    // 获取最大的元素
    int getMin(vector<ListNode*>& lists)
    {
        if (!lists.size())
        {
            return -1;
        }
        int index = -1;
        int value = 0;
        for (int i = 0; i<lists.size();i++)
        {
            if (!lists[i]) {
                continue;
            }

            if (index == -1)
            {
               index = i;
               value = lists[index]->val;
               continue;
            }

            if (lists[i]->val < value)
            {
               index = i;
               value = lists[i]->val;
            }
        }
        return index;
    }

    // 更新指定索引的头节点
    void Update(vector<ListNode*>& lists, int index)
    {
        if (index != -1)
        {
            lists[index] = lists[index]->next;
        }
    }

    // 合并节点
    ListNode* mergeKLists(vector<ListNode*>& lists) {
        ListNode* dummyNode = new ListNode(0, nullptr);
        ListNode* result = dummyNode;
        while (getNull(lists))
        {
            int index = getMin(lists);
            if (index != -1)
            {
                dummyNode->next = lists[index];
                Update(lists, index);
                dummyNode = dummyNode->next;
            }
        }
        return result->next;
    }
};