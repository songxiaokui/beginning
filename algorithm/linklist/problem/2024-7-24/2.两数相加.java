/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode head = new ListNode();
        if (l1 == null && l2 == null) {
            return null;
        }
        ListNode current = head;
        int carry = 0;
        while (l1 != null || l2 != null) {
            int v1 = 0;
            int v2 = 0;
            ListNode node = new ListNode();
            if (l1 != null) {
                v1 = l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                v2 = l2.val;
                l2 = l2.next;
            }
            node.val = (v1+v2+carry) % 10;
            carry = (v1+v2+carry)/10;
            current.next = node;
            current = current.next;

        }
        if (carry != 0) {
            current.next = new ListNode(carry);
        }
        return head.next;
    }
}