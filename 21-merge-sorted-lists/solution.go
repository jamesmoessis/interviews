package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current1 := list1
	current2 := list2
	var current *ListNode
	if current1 == nil && current2 == nil {
		return nil
	}
	var head *ListNode

	for {
		if head == nil && current != nil {
			head = current
		}
		if current1 == nil && current2 == nil {
			break
		}
		if current1 == nil {
			// current2 is not nil
			current = current2
			break
		}
		if current2 == nil {
			// current1 is not nil
			current = current1
			break
		}

		// both aren't nil
		if current1.Val < current2.Val {
			current.Next = current1
		} else {
			current.Next = current2
		}
	}
	return head
}
