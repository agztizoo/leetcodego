package medium

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, tail *ListNode
    v, t := 0, 0
    for l1 != nil || l2 != nil {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        v, t = (v1+v2+t)%10, (v1+v2+t)/10
        if head == nil {
            head = &ListNode{
                Val: v,
            }
            tail = head
        } else {
            n := &ListNode{
                Val: v,
            }
            tail.Next = n
            tail = n
        }
    }
    if t > 0 {
        n := &ListNode{
            Val: t,
        }
        tail.Next = n
    }
    return head
}
