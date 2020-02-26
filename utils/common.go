package utils

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
    node := &ListNode {Val: val, Next: next}
    return node
}