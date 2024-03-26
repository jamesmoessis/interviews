package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	current1 := list1
	current2 := list2
	if current1 == nil && current2 == nil {
		return nil
	}
	// "Blank" head to use just so there's something to attach to
	// Returned head will be head.Next
	head := &ListNode{}
	current := head

	for !(current1 == nil && current2 == nil) {
		if current1 == nil {
			// current2 is not nil
			current.Next = current2
			break
		}
		if current2 == nil {
			// current1 is not nil
			current.Next = current1
			break
		}

		// both aren't nil
		if current1.Val < current2.Val {
			current.Next = current1
			current1 = current1.Next
		} else {
			current.Next = current2
			current2 = current2.Next
		}
		current = current.Next
	}
	return head.Next
}
