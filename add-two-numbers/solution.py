# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = result = ListNode()
        carry = 0
        _sum = 0
        while True:
            if l1 or l2 or carry:
                result.next = ListNode()
                result = result.next
                _sum = carry
                carry = 0
                if l1 and l2:
                    _sum += l1.val + l2.val
                    result.val = _sum % 10
                    carry = _sum // 10
                    l1 = l1.next
                    l2 = l2.next
                elif l1:
                    _sum += l1.val
                    result.val = _sum % 10
                    l1 = l1.next
                    carry = _sum // 10
                elif l2:
                    _sum += l2.val
                    l2 = l2.next
                    result.val = _sum % 10
                    carry = _sum // 10
                elif _sum:
                    result.val = _sum
            else:
                break
        return head.next
