# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        result = ListNode(0, None)
        current = result
        if l1 is None and l2 is None:
            return result.next
        carry = 0
        while (l1 or l2):
            # 计算值 创建节点 移动两个指针
            v1 = 0
            v2 = 0
            val = 0
            if l1:
                v1 = l1.val
                l1 = l1.next
            if l2:
                v2 = l2.val
                l2 = l2.next
            node = ListNode((v1+v2+carry)%10, None)
            current.next = node
            current = current.next
            carry = (v1+v2+carry)/10

        # 处理最后进位
        if carry != 0:
            current.next = ListNode(carry, None)
        return result.next