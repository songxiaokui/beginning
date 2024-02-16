/*
19. 删除链表的倒数第 N 个结点
中等
相关标签
相关企业
提示
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：
输入：head = [1], n = 1
输出：[]
*/

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        // 解题思路
        // 首先可以定义一个哑节点 方便处理头部元素
        ListNode* dummyNode = new ListNode(0, head);
        // 删除元素 通常要找删除该元素的前驱节点和后继节点 然后让前驱节点的 next 指向被删除节点的后继节点即可
        // 使用快慢指针用来找指定位置的元素
        // 如找倒数第 n 个
        // 可以先让快节点走 n步 然后慢节点和快节点同时走 直到走到尾部 就可以找到删除节点的前驱节点
        // head--> 1 --> 2 --> 3 --> 4 --> 5
        // 当前 fast 和 slow 在 head
        // 删除倒数 2 号元素 也就是 4
        // 快节点走 2 步 到达节点 2 号位置
        // 此时慢节点和快节点同时走 直到快节点到尾部
        // 也就是走 3步到尾部
        // 此时慢节点刚好走到删除元素 的前驱节点 3，fast 节点为 nullptr
        ListNode* fast = dummyNode;
        ListNode* slow = dummyNode;
        int i = 0;
        while (i < n+1)
        {
            fast = fast->next;
            i++;
        }

        // 然后快慢指针同时走
        while (fast)
        {
            fast = fast->next;
            slow = slow->next;
        }

        // 找到被删除的节点
        ListNode* deleteNode = slow->next;
        // 删除节点
        slow->next = deleteNode->next;
        // 释放删除的节点内存
        delete deleteNode;

        return dummyNode->next;
    }
};
