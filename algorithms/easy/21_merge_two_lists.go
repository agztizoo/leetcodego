package easy

import (
	. "github.com/leetcodego/utils"
)

// MergeTwoLists 21. Merge Two Sorted Lists
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := NewListNode(0, nil)
    curr := head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next, l1 = l1, l1.Next
        } else {
            curr.Next, l2 = l2, l2.Next
        }
        curr = curr.Next
    }
    if l1 == nil {
        curr.Next = l2
    }
    if l2 == nil {
        curr.Next = l1
    }
    return head.Next
}
