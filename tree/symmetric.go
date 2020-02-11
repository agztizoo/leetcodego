package tree

func isSymmetric(root *TreeNode) bool {
    return (root == nil) || isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
    if (p == nil && q == nil) {
        return true;
    }
    if (p == nil || q == nil) {
        return false;
    }
    sameLeft := isSymmetricTree(p.Left, q.Right)
    sameRight := isSymmetricTree(p.Right, q.Left)
    return (p.Val == q.Val) && sameLeft && sameRight
}
