package tree

import "github.com/leetcodego/utils"

func isSymmetric(root *utils.TreeNode) bool {
    return (root == nil) || isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *utils.TreeNode, q *utils.TreeNode) bool {
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
