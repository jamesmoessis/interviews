"""
Given the head of a linked list, remove the nth node from the end of the list and return its head.

1st attempt: 
find length of list 
write function to remove(index)

call remove(len(l) - n)
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def length(head: ListNode) -> int:
    curnt = head
    i = 0
    while curnt:
        i += 1
        curnt = curnt.next
    return i

def remove(head: ListNode, index: int) -> ListNode:
    current = head
    previous = None
    if index == 0:
        return head.next
    for i in range(index):
        if i == index:
            break
        previous = current
        current = current.next
    previous.next = current.next
    return head
    

class Solution:
    def removeNthFromEnd(self, head: ListNode, n: int) -> ListNode:
        if head == None:
            return None
        l = length(head)
        return remove(head, l-n)


def generate_list(l):
    head = node = ListNode(l[0])
    for num in l[1:]:
        node = ListNode(num)
        head.next = node
    return head


if __name__ == '__main__':
    l = generate_list([1,2,3,4,5])
    s = Solution().removeNthFromEnd(l, 2)
    print(s.val)